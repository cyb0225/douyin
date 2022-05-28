# 数据库 database（repositroy）

最初实现思路：深思熟路后，拟定最初方案，创建多个表，避免数据表之间复杂关系。且采用单一mysql形式。

https://blog.csdn.net/weixin_42405670/article/details/104219333

**User数据库需要存储的数据**

- id （用户id）（主键，登录后用于用户数据查询）
- username （用户名）
- password （用户密码）(加密后存储到这里)
- followcount （关注的总数）
- followercount （被关关注的总数）
**Video数据库需要存储的数据**
- id （视频id） （主键）
- user (auther 发布者的ID)
- title (视频标题)
- play_url 	(视频播放地址)
- cover_url 	(视频封面地址)
- favorite_count   (视频的点赞总数)
- comment_count  (视频的评论总数)

**follow 关注以及被关注**

-  follow_id (关注者的id)
- follower_id （被关注者的id)
- status (是否关注的状态)

**favourite**

- video_id （视频id）
- user_id (用户id)
- status (是否喜欢的状态)

**comment**
- user_id 
- video_id