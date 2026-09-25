# Age Tracker API

This is a GOLANG application which can be used to track Age of Users. This API uses Golang, Gin and PostgreSQL DB.

It has the following endpoint:

### Get List of People
### Add new People

## Deployment

The API is hosted on [Render](https://render.com) (free web service) using the `render.yaml` Blueprint, backed by a free Postgres database on [Neon](https://neon.tech).

1. Create a Neon project and run `DbSchema/schema.sql` then `DbSchema/stored_procedure.sql` in its SQL editor.
2. In Render, choose **New > Blueprint**, select this repository, and paste the Neon connection string as `DATABASE_URL` when prompted.
3. Every push to `main` redeploys automatically.

Free Render services sleep after 15 minutes without traffic, so the first request after a pause can take up to a minute.
