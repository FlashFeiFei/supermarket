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


--微信客户表
CREATE TABLE `supermarket_wechat_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `openid` varchar(30) NOT NULL COMMENT '微信openid',
  `unionid` varchar(30) NOT NULL COMMENT '微信unionid',
  `account_type` int(11) NOT NULL COMMENT '微信账号类型、小程序、公众号等',
  `user_id` int(11) NOT NULL COMMENT 'supermarket的id',
  `createtime` int(11) NOT NULL COMMENT '创建时间',
  `updatetime` int(11) NOT NULL COMMENT '更新时间',
  `deletetime` int(11) NOT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='微信用户表';


-- 附件表
CREATE TABLE `supermarket_attachment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) DEFAULT NULL COMMENT '源文件名',
  `filepath` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `filetype` int(11) NOT NULL DEFAULT '0' COMMENT '文件类型，图片、视频、音频',
  `links` varchar(255) DEFAULT NULL COMMENT '外部的链接',
  `updatetime` int(11) NOT NULL DEFAULT '0',
  `createtime` int(11) NOT NULL DEFAULT '0',
  `deletetime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片、视频等资源。附件表';


-- 轮播图
CREATE TABLE `supermarket_banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `attachment_id` int(11) NOT NULL COMMENT '附件id',
  `title` varchar(255) NOT NULL COMMENT '轮播图标题',
  `createtime` int(11) NOT NULL,
  `updatetime` int(11) NOT NULL,
  `deletetime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='轮播图';


-- 商品分类
CREATE TABLE `supermarket_goods_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '类名',
  `createtime` int(11) NOT NULL,
  `updatetime` int(11) NOT NULL,
  `deletetime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品分类';


-- 商品
CREATE TABLE `supermarket_goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '商品名称',
  `details` varchar(255) NOT NULL COMMENT '商品详情',
  `parameter` varchar(255) NOT NULL COMMENT '商品参数。比如多少多少克',
  `price` int(10) NOT NULL COMMENT '单价，单位分',
  `attachment_id` int(11) NOT NULL COMMENT '附件id',
  `theme_id` int(11) NOT NULL DEFAULT '0' COMMENT '主题id',
  `category_id` int(11) NOT NULL COMMENT '分类id',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '推荐程度',
  `inventory` tinyint(11) NOT NULL COMMENT '库存，1表示有，0表示',
  `createtime` int(11) NOT NULL,
  `updatetime` int(11) NOT NULL,
  `deletetime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表';


-- 主题表
CREATE TABLE `supermarket_theme` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `attachment_id` int(11) NOT NULL COMMENT '附件id',
  `title` varchar(255) NOT NULL COMMENT '主题名',
  `createtime` int(11) NOT NULL,
  `updatetime` int(11) NOT NULL,
  `deletetime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='主题表';





