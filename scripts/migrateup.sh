#!/biset -e

if [ -f .env ]; then
    source .env
fi

echo "Running database migrations..."
cd sql/schema
goose turso $DATABASE_URL up
echo "Migrations completed successfully!"
