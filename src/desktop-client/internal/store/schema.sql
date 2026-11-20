CREATE TABLE IF NOT EXISTS "projects" (
                                          "id" TEXT NOT NULL UNIQUE,
                                          "name" VARCHAR,
                                          "directory" VARCHAR,
                                          "git-local" VARCHAR,
                                          "git-online" VARCHAR,
                                          "created" TIMESTAMP,
                                          "modified" TIMESTAMP,
                                          PRIMARY KEY("id"),
    FOREIGN KEY ("id") REFERENCES "sessions"("project")
    ON UPDATE NO ACTION ON DELETE NO ACTION
    );

CREATE TABLE IF NOT EXISTS "t1-heartbeats" (
                                               "id" TEXT NOT NULL,
                                               "session" TEXT,
                                               "relative-path" VARCHAR,
                                               "timestamp" TIMESTAMP,
                                               PRIMARY KEY("id"),
    FOREIGN KEY ("session") REFERENCES "sessions"("id")
    ON UPDATE NO ACTION ON DELETE NO ACTION
    );

CREATE TABLE IF NOT EXISTS "sessions" (
                                          "id" TEXT NOT NULL,
                                          "started" TIMESTAMP,
                                          "ended" TIMESTAMP,
                                          "project" TEXT,
                                          PRIMARY KEY("id")
    );
