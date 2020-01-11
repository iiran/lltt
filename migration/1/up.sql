--- users beg
CREATE TABLE public.users (
  id BIGSERIAL PRIMARY KEY,
  username CHARACTER VARYING(128) NOT NULL UNIQUE,
  display_name CHARACTER VARYING(128),
  email CHARACTER VARYING(128),
  PASSWORD CHARACTER VARYING(256),
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW() NOT NULL,
  disable_at TIMESTAMP WITHOUT TIME ZONE DEFAULT '1970-01-01 00:00:00' :: timestamp without time zone NOT NULL,
  score INTEGER DEFAULT 0 NOT NULL,
  weight INTEGER DEFAULT 0 NOT NULL,
  gender SMALLINT NOT NULL DEFAULT -1,
  birth TIMESTAMP WITHOUT TIME ZONE
);

CREATE INDEX username_index ON public.users (username);

COMMENT ON COLUMN users.gender IS '-1: unknown, 0: woman, 1:man';

--- users end
--- post_change beg
CREATE TABLE post_change (
  id BIGSERIAL PRIMARY KEY,
  post_id BIGINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  change_type SMALLINT NOT NULL,
  old TEXT NOT NULL,
  new TEXT NOT NULL
);

CREATE INDEX post_id_index ON public.post_change(post_id);

CREATE INDEX change_name_index ON public.post_change(change_type);

COMMENT ON COLUMN post_change.change_type IS '1:title, 2:content, 3: describe, 4: tag_name, 5: author';

--- post_change end
--- post_tag beg
CREATE TABLE post_tag (
  name varchar(256) NOT NULL PRIMARY KEY,
  created_at timestamp NOT NULL DEFAULT NOW(),
  disable_at timestamp NOT NULL DEFAULT '1970-01-01 00:00:00'
);

--- post_tag end
--- posts beg
CREATE TABLE posts (
  id bigserial PRIMARY KEY,
  title varchar(512) NOT NULL,
  DESCRIBE text NOT NULL,
  content text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  disable_at timestamp NOT NULL DEFAULT '1970-01-01 00:00:00',
  creator bigint NOT NULL,
  author bigint NOT NULL,
  tag_name varchar(256)
);

CREATE INDEX creator_index ON public.posts(creator);

CREATE INDEX tag_name_index ON public.posts(tag_name);

--- posts end
--- reply beg
CREATE TABLE reply (
  id bigserial PRIMARY KEY,
  content text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  disable_at timestamp NOT NULL DEFAULT '1970-01-01 00:00:00',
  author bigint NOT NULL,
  reply_type SMALLINT NOT NULL,
  reply_to bigint NOT NULL
);

CREATE INDEX reply_type_index ON public.reply(reply_type);

CREATE INDEX reply_to_index ON public.reply(reply_to);

COMMENT ON COLUMN reply.reply_type IS '1:user, 2:post, 3: reply';

--- reply end
--- message beg
CREATE TABLE short_message (
  id bigserial PRIMARY KEY,
  content text NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  disable_at timestamp NOT NULL DEFAULT '1970-01-01 00:00:00',
  author bigint NOT NULL,
  message_type smallint NOT NULL,
  message_to bigint NOT NULL,
  is_withdraw SMALLINT NOT NULL DEFAULT 0
);

CREATE INDEX message_type_index ON public.short_message(message_type);

CREATE INDEX message_author_index ON public.short_message(author);

CREATE INDEX message_to_index ON public.short_message(message_to);

COMMENT ON COLUMN short_message.message_type IS '1: user, 2: group, 3: system';

--- message end
--- group start
CREATE TABLE groups (
  id bigserial PRIMARY KEY,
  name varchar(128) NOT NULL,
  describe varchar(1024) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  disable_at timestamp NOT NULL DEFAULT '1970-01-01 00:00:00',
  weight int NOT NULL,
  score int NOT NULL
);

create index groups_name_index on public.groups(name);

create index groups_weight_index on public.groups(weight);
---group end

insert into users (username) values ('yiranfeng');