This directory contains a storage factory interface that can be used to create different storage backends

The interface simply offers CRUD operations on blobs/files that can be translated into local storage, cloud storage, key value database or relational database calls.

The interface is more akin to a nosql database / document store than a relational database so I would probably change or augment this interface with something to handle relationships between objects.

I've decided to go with PostgreSQL as it offers relational and JSON document style storage. It's commonly used and understood and cheap.

```sql
CREATE DATABASE aosdesk
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
```

```sql
CREATE ROLE aosdesk WITH
	LOGIN
	NOSUPERUSER
	NOCREATEDB
	NOCREATEROLE
	INHERIT
	NOREPLICATION
	CONNECTION LIMIT -1
	PASSWORD 'xxxxxx';
COMMENT ON ROLE aosdesk IS 'Generic user account for web backend';
```
