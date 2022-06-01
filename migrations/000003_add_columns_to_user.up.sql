alter table user add column `finish_problem_num` int(11) NOT NULL DEFAULT 0 COMMENT '完成题目数量' after `email`;
alter table user add column `submit_num` int(11) NOT NULL DEFAULT 0 COMMENT '提交题目数量' after `finish_problem_num`;
