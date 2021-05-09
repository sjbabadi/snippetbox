CREATE TABLE snippets (
  id SERIAL PRIMARY KEY,
  title varchar NOT NULL,
  content varchar NOT NULL,
  created_at timestamp NOT NULL,
  expires_at timestamp NOT NULL
);

CREATE INDEX idx_snippets_created_at ON snippets(created_at);

INSERT INTO snippets (title, content, created_at, expires_at) VALUES (
  'An old silent pond',
  'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
  now(),
  now() + INTERVAL '7 days'
);

INSERT INTO snippets (title, content, created_at, expires_at) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    now(),
    now() + INTERVAL '7 days'
);

INSERT INTO snippets (title, content, created_at, expires_at) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    now(),
    now() + INTERVAL '7 days'
);