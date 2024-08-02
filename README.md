# Heimdall

Heimdall is a Golang service that interacts with GitHub's public API to fetch repository and commit data, store it in a
PostgreSQL database, and continuously monitor repositories for changes.

## Features

- Fetches repository metadata and commit history from GitHub.
- Stores data in a PostgreSQL database.
- Periodically checks for new commits and updates.
- Provides API to query stored data.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or later
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [GitHub API Token](https://github.com/settings/tokens)

### Installation

1. **Clone the repository:**

```shell
git clone https://github.com/isongjosiah/heimdall.git
cd heimdall

```

2. **Set up environment variables:**
   Create a `.env` file in the project root and add the following

```dotenv
GITHUB_TOKEN=your_github_token
DATABASE_URL=postgresql://heimdall:q7ilD9qSKeLgJqRWW8K6CVbZNuELPcFT@dpg-cqll7h3v2p9s73b1vc20-a.frankfurt-postgres.render.com/heimdalldb=
RMQ_URL=amqps://qlaajerz:2lKOabiVLRUokblIXGjWkuuEqcgKqxmC@seal.lmq.cloudamqp.com/qlaajerz

```

3. **Build the Docker image:**

```shell
docker build -t heimdall:latest .

```

4. **Run the application:**

```shell
docker run --env-file=.env heimdall:latest

```

### Configuration

Configuration settings are managed using environment variables. You can
customize
settings such as database connection strings and API tokens directly via environment
variables.

### Usage

- **Fetch Repository Data:**
    - The application automatically fetches data for specified repositories.
- **Monitor for Changes:**
    - The application checks for updates at regular intervals (default: every hour)

### Database Schema

The application uses two main tables:

1. **Repositories**
    - Stores metadata about each repository.
   
      | Column Name       | PostgreSQL Type | Description                                         |
      |-------------------|-----------------|-----------------------------------------------------|
      | id                | UUID            | Primary key identifier for the repository.          |
      | name              | TEXT            | Name of the repository.                             |
      | owner             | TEXT            | Owner of the repository.                            |
      | description       | TEXT            | Description of the repository.                      |
      | url               | TEXT            | URL link to the repository.                         |
      | language          | TEXT            | Primary programming language of the repository.     |
      | fork_count        | INTEGER         | Number of forks of the repository.                  |
      | stars_count       | INTEGER         | Number of stars the repository has received.        |
      | open_issue_count  | INTEGER         | Number of open issues in the repository.            |
      | watchers_count    | INTEGER         | Number of watchers of the repository.               |
      | pull_from         | TIMESTAMP       | Date and time from which to pull updates.           |
      | initial_pull_done | BOOLEAN         | Flag indicating if the initial pull was completed.  |
      | created_at        | TIMESTAMP       | Date and time the repository was created.           |
      | added_at          | TIMESTAMP       | Date and time the repository was added.             |
      | updated_at        | TIMESTAMP       | Date and time the repository was last updated.      |

2. **Commits**
    - Stores commit information for each repository.

      | Column Name | PostgreSQL Type | Description                              |
      |-------------|-----------------|------------------------------------------|
      | id          | UUID            | Primary key identifier for the commit.   |
      | repo_id     | UUID            | Unique identifier for the repository.    |
      | sha         | TEXT            | Unique SHA string for the commit.        |
      | message     | TEXT            | Commit message.                          |
      | author      | JSONB           | JSON object containing author details.   |
      | commit_date | TIMESTAMP       | Date and time of the commit.             |
      | url         | TEXT            | URL link to the commit.                  |
      | added_at    | TIMESTAMP       | Date and time the commit was added.      |

    - Notes
        - `id`: Primary key, stored as a UUID for uniqueness.
        - `repo_id` and `sha` are combined to form a unique constraint (`unique:repo_commit`).
        - `author`: Stored as a JSONB to accommodate complex author details.

### API Endpoints

- **Add Repository To Track:**
    - **Endpoint:** `/repositories`
    - Adds a GitHub repository to track

```json
{
  "repo_name": "chromium",
  "repo_owner": "chromium",
  "fetch_commit_from": "2024-01-01T00:00:00Z"
}
```

- **Get Top N Commit Authors:**
    - **Endpoint:** `/commits/top-contributors`
    - Retrieves the top N commit authors by commit counts
- **Get Commits by Repository:**
    - **Endpoint:** `repositories/commits?repo-name={repoName}`
    - Retrieves paginated commits for a given repository.

API Endpoint documentation is available [here](https://documenter.getpostman.com/view/10427889/2sA3rwLDeg)

Deployed version of the application is available [here](https://heimdall-8hge.onrender.com)

### Running Tests

Run unit tests using the Go testing framework:

```shell
go test -v ./...
```

### Deployment

This application can be deployed using Docker and Kubernetes. A sample deployment configuration (`deployment.yaml`) is
provided for Kubernetes.

1. **Build and push the Docker image:**

```shell
docker build -t your-docker-repo/heimdall:latest .
docker push your-docker-repo/heimdall:latest

```

2. **Deploy to Kubernetes:**

```shell
kubectl apply -f deployment.yaml

```

### Optimizing Data Retrieval

to optimize data retrieval and operations, two indexes were created on the `commits` table:

- **Index on `repo_id`**:
    - This significantly reduced the time required for operations involving repository-specific data, such as resetting
      a repository's commit collection or retrieve collections for a repository.
    - **Benefit**: The time to delete records when resetting a repo collection decreased from 2 minutes to 1.5 seconds
    ```sql
    CREATE INDEX idx_commits_repo_id ON commits(repo_id)
    ```

- **Index on `author`**:
    - This index optimizes queries that filter or sort commits by author, improving performance of operations like
      retrieving
      the top N commit authors.
  ```sql
  CREATE INDEX idx_commits_author ON commits (author);
  ```
