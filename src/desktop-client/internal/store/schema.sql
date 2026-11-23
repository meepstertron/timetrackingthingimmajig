CREATE TABLE IF NOT EXISTS "projects" (
                                          "id" TEXT NOT NULL UNIQUE,
                                          "name" VARCHAR,
                                          "directory" VARCHAR,
                                          "git-local" VARCHAR,
                                          "git-online" VARCHAR,
                                          "created" TIMESTAMP,
                                          "modified" TIMESTAMP,
                                          PRIMARY KEY("id")
    );

CREATE TABLE IF NOT EXISTS "sessions" (
                                          "id" TEXT NOT NULL,
                                          "started" TIMESTAMP,
                                          "ended" TIMESTAMP,
                                          "project" TEXT,
                                          PRIMARY KEY("id"),
    FOREIGN KEY ("project") REFERENCES "projects"("id")
    ON UPDATE NO ACTION ON DELETE NO ACTION
    );

CREATE TABLE IF NOT EXISTS "t1-heartbeats" (
                                               "id" TEXT NOT NULL,
                                               "session" TEXT,
                                               "relative-path" VARCHAR,
                                               "timestamp" TIMESTAMP,
                                               "action" VARCHAR, -- edit, create, delete, rename, new-folder
                                               "lines-added" INTEGER NOT NULL DEFAULT 0,
                                               "lines-removed" INTEGER NOT NULL DEFAULT 0,
                                               "hunks" VARCHAR,
                                               "file-hash" VARCHAR,
                                               PRIMARY KEY("id"),
    FOREIGN KEY ("session") REFERENCES "sessions"("id")
    ON UPDATE NO ACTION ON DELETE NO ACTION
    );