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
settings such as database connection strings and API tokens  directly via environment
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

| Column | Type    |
|--------|---------|
| ID     | VARCHAR |
| Name   | VARCHAR |

2. **Commits**
    - Stores commit information for each repository.
  
| Column | Type    |
|--------|---------|
| ID     | VARCHAR |
| Name   | VARCHAR |

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