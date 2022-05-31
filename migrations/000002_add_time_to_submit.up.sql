alter table submit add column `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP;
alter table submit add column `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP;
alter table submit add column `deleted_at` datetime NULL DEFAULT null;