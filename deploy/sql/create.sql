
CREATE TABLE IF NOT EXISTS users (
    id              CHAR(32) PRIMARY KEY,
    username        VARCHAR(127) NOT NULL,
    password        VARCHAR(127) NOT NULL,
    nick_name       VARCHAR(127) NOT NULL,
    display_name    VARCHAR(127) NOT NULL,
    gender          CHAR NOT NULL DEFAULT 0,  -- M:男性 F:女性 O:其他
    birthdate       DATE,                     -- 出生日期
    phone           VARCHAR(20) NULL,
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE UNIQUE INDEX IF NOT EXISTS uk_users_user_pool_username ON users(username);
CREATE UNIQUE INDEX IF NOT EXISTS uk_users_phone ON users(phone);
