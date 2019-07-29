module.exports = {
    "type": "mysql",
    "host": process.env["MYSQL_HOST"],
    "port": process.env["MYSQL_PORT"],
    "username": process.env["MYSQL_USERNAME"],
    "password": process.env["MYSQL_PASSWORD"],
    "database": "college",
    "synchronize": true,
    "logging": false,
    "migrationsTableName": "custom_migration_table",
    "entities": [
        "src/entity/**/*.ts"
    ],
    "migrationsRun": true,
    "migrations": [
        "src/migration/**/*.ts"
    ],
    "subscribers": [
        "src/subscriber/**/*.ts"
    ],
    "cli": {
        "entitiesDir": "src/entity",
        "migrationsDir": "src/migration",
        "subscribersDir": "src/subscriber"
    }
}