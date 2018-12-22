-- 客户表

CREATE TABLE `supermarket_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) NOT NULL DEFAULT '' COMMENT '登录账号',
  `password` varchar(40) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(10) NOT NULL DEFAULT '' COMMENT '昵称',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '1正常用户，0被封',
  `email` varchar(32) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `createtime` int(11) NOT NULL COMMENT '创建时间',
  `updatetime` int(11) NOT NULL COMMENT '更新时间',
  `deletetime` int(11) NOT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='客户表';