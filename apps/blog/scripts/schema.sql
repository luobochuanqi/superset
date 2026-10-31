-- utf8mb4 才能存中文和 emoji；InnoDB 支持事务和外键

USE blog;

-- 标签表
CREATE TABLE blog_tag (
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增',
    name       VARCHAR(64)  NOT NULL COMMENT '标签名',
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME     NULL DEFAULT NULL COMMENT '软删除：非 NULL 即视为已删除（GORM 约定）',
    UNIQUE KEY uk_name (name) COMMENT '标签名不能重复，这也是一个索引，按名字查很快'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签表';

-- 文章表
CREATE TABLE blog_article (
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键，自增',
    title           VARCHAR(256) NOT NULL COMMENT '标题',
    content         TEXT         NOT NULL COMMENT '正文',
    cover_image_url VARCHAR(512) NOT NULL DEFAULT '' COMMENT '封面图',
    created_at      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at      DATETIME     NULL DEFAULT NULL COMMENT '软删除：非 NULL 即视为已删除（GORM 约定）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

-- 文章-标签关联表：一篇文章多个标签、一个标签多篇文章（多对多），必须用中间表
CREATE TABLE blog_article_tag (
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    article_id BIGINT UNSIGNED NOT NULL COMMENT '属于哪篇文章',
    tag_id     BIGINT UNSIGNED NOT NULL COMMENT '属于哪个标签',
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at DATETIME     NULL DEFAULT NULL COMMENT '软删除：非 NULL 即视为已删除（GORM 约定）',
    UNIQUE KEY uk_article_tag (article_id, tag_id) COMMENT '同一篇文章不能挂同一个标签两次',
    KEY idx_tag_id (tag_id) COMMENT '按标签反查文章要快',
    CONSTRAINT fk_at_article FOREIGN KEY (article_id) REFERENCES blog_article (id) ON DELETE CASCADE,
    CONSTRAINT fk_at_tag     FOREIGN KEY (tag_id)     REFERENCES blog_tag (id)     ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章-标签关联表';
