Gator CLI
Gator is a command-line interface (CLI) tool for managing RSS feeds and user subscriptions. It allows users to register, log in, add and follow RSS feeds, and aggregate content from specified feed URLs. The application uses PostgreSQL for data storage and is built with Go.
Prerequisites
To run the Gator CLI, you need the following installed:
PostgreSQL (version 13 or later recommended):
Install PostgreSQL: Official PostgreSQL Downloads

Ensure the PostgreSQL server is running and you have a user with database creation privileges.

Create a database named gator:
bash

psql -U <your-username> -c "CREATE DATABASE gator;"

Go (version 1.20 or later):
Install Go: Official Go Downloads

Verify installation:
bash

go version

Installation
To install the Gator CLI, use the go install command:
bash

go install github.com/your-username/gator@latest

Replace your-username/gator with the actual repository path if different. This installs the gator binary to $GOPATH/bin (or $HOME/go/bin by default). Ensure $GOPATH/bin is in your $PATH:
bash

export PATH=$PATH:$HOME/go/bin

Verify installation:
bash

gator --version

Dependencies
The CLI uses external tools and libraries:
Goose: For database migrations (installed via go get github.com/pressly/goose/v3).

SQLC: For generating Go code from SQL queries (install below).

PostgreSQL driver: github.com/lib/pq (included in go.mod).

Install SQLC:
bash

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

Setup
1. Clone the Repository
If you’re contributing or running from source, clone the repository and navigate to it:
bash

git clone https://github.com/your-username/gator.git
cd gator

2. Set Up the Configuration File
Gator requires a configuration file at ~/.gatorconfig.json. Create it with the following content:
json

{
  "db_url": "postgres://<username>:<password>@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}

Replace <username> and <password> with your PostgreSQL credentials.

Ensure the db_url points to the gator database.

The current_user_name is updated when you log in or register.

3. Apply Database Migrations
Run Goose migrations to create the necessary tables (users, feeds, feed_follows):
bash

goose -dir sql/schema postgres "<your_db_url>" up

Replace <your_db_url> with the same db_url from ~/.gatorconfig.json.
4. Generate SQLC Code
Generate Go code for database queries:
bash

sqlc generate

5. Run the CLI
Run the CLI from the repository root (if cloned) or use the installed binary:
bash

go run . <command>
# OR
gator <command>

Available Commands
Here are some key commands you can use with the Gator CLI:
Register a User:
Creates a new user and sets them as the current user.
bash

gator register <username>

Example:
bash

gator register lane

Output: user created: lane

Log In:
Sets an existing user as the current user.
bash

gator login <username>

Example:
bash

gator login lane

Output: logged in as lane

Add a Feed:
Adds a new RSS feed and automatically follows it for the current user (requires login).
bash

gator addfeed "<feed_name>" "<feed_url>"

Example:
bash

gator addfeed "Lane's Blog" "https://www.wagslane.dev/index.xml"

Output: Feed details and User lane is now following feed Lane's Blog

Follow a Feed:
Follows an existing feed by URL (requires login).
bash

gator follow "<feed_url>"

Example:
bash

gator follow "https://www.wagslane.dev/index.xml"

Output: User lane is now following feed Lane's Blog

Unfollow a Feed:
Unfollows a feed by URL (requires login).
bash

gator unfollow "<feed_url>"

Example:
bash

gator unfollow "https://www.wagslane.dev/index.xml"

Output: User lane has unfollowed feed Lane's Blog

List Followed Feeds:
Lists all feeds the current user is following (requires login).
bash

gator following

Example:
bash

gator following

Output: * Lane's Blog

List All Feeds:
Lists all feeds in the database with their names, URLs, and creator usernames.
bash

gator feeds

Example:
bash

gator feeds

Output:

Name: Lane's Blog
URL: https://www.wagslane.dev/index.xml
User: lane

Aggregate a Feed:
Fetches and prints the content of a specific RSS feed (for testing feed parsing).
bash

gator agg

Example:
bash

gator agg

Output: Structured RSS feed data from https://www.wagslane.dev/index.xml.

Development
To contribute or modify the CLI:
Track changes with Git:
bash

git add .
git commit -m "Your commit message"
git push origin main

Install dependencies:
bash

go mod tidy

Run tests:
bash

go test ./...

Troubleshooting
Database Connection Errors:
Verify ~/.gatorconfig.json has the correct db_url.

Ensure PostgreSQL is running and the gator database exists.

SQLC Errors:
Run sqlc generate after modifying queries.

Check sqlc.yaml for correct schema and query paths.

Command Failures:
Ensure you’re logged in for commands like addfeed, follow, unfollow, and following.

Check error messages for specific issues (e.g., error: feed url already exists).

For issues, open a GitHub issue or check the PostgreSQL and Go documentation.

