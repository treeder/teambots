# Teamster

The ultimate GitHub/Slack app for dev teams. Powered by Fn.

## Deploying Teamster

### 1. Get a SQL database setup somewhere

There's free and easy ones at Heroku.

Setup your app:

```sh
# Export the Fn API URL you're using:
export FN_API_URL=https://fn-do.funcy.run
# Create app
fn apps create teamster
# Set your database credentials URL
fn apps config set teamster DB_URL postgres://XYZ
# Get goose and goose up
goose up `fn apps config get DB_URL`
# Deploy the app
fn deploy --all --app teamster
```
