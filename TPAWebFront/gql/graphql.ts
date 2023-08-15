/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
};

export type Comment = {
  __typename?: 'Comment';
  comments?: Maybe<Array<Maybe<Comment>>>;
  content: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  likeCount: Scalars['Int']['output'];
  liked?: Maybe<Scalars['Boolean']['output']>;
  likes?: Maybe<Array<Maybe<CommentLike>>>;
  parentComment?: Maybe<Comment>;
  parentPost?: Maybe<Post>;
  replyCount: Scalars['Int']['output'];
  user: User;
};

export type CommentLike = {
  __typename?: 'CommentLike';
  commentId: Scalars['ID']['output'];
  user: User;
};

export type Conversation = {
  __typename?: 'Conversation';
  group?: Maybe<Group>;
  id: Scalars['ID']['output'];
  messages?: Maybe<Array<Maybe<Message>>>;
  users: Array<ConversationUsers>;
};

export type ConversationUsers = {
  __typename?: 'ConversationUsers';
  conversationId: Scalars['ID']['output'];
  user: User;
};

export type Friend = {
  __typename?: 'Friend';
  accepted: Scalars['Boolean']['output'];
  receiver: User;
  sender: User;
};

export type FriendInput = {
  receiver: Scalars['ID']['input'];
  sender: Scalars['ID']['input'];
};

export type Group = {
  __typename?: 'Group';
  about: Scalars['String']['output'];
  background: Scalars['String']['output'];
  chat?: Maybe<Conversation>;
  createdAt: Scalars['Time']['output'];
  files?: Maybe<Array<Maybe<GroupFile>>>;
  id: Scalars['ID']['output'];
  isAdmin: Scalars['Boolean']['output'];
  joined: Scalars['String']['output'];
  memberCount: Scalars['Int']['output'];
  members: Array<Member>;
  name: Scalars['String']['output'];
  posts?: Maybe<Array<Maybe<Post>>>;
  privacy: Scalars['String']['output'];
};

export type GroupFile = {
  __typename?: 'GroupFile';
  groupID: Scalars['ID']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  type: Scalars['String']['output'];
  uploadedAt: Scalars['Time']['output'];
  uploadedBy: User;
  url: Scalars['String']['output'];
};

export type Member = {
  __typename?: 'Member';
  approved: Scalars['Boolean']['output'];
  groupId: Scalars['ID']['output'];
  role: Scalars['String']['output'];
  user: User;
};

export type Message = {
  __typename?: 'Message';
  conversationId: Scalars['ID']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  image?: Maybe<Scalars['String']['output']>;
  message?: Maybe<Scalars['String']['output']>;
  post?: Maybe<Post>;
  sender: User;
};

export type Mutation = {
  __typename?: 'Mutation';
  acceptFriend?: Maybe<Friend>;
  activateUser: User;
  addFriend?: Maybe<Friend>;
  authenticateUser: Scalars['String']['output'];
  createComment?: Maybe<Comment>;
  createConversation?: Maybe<Conversation>;
  createGroup: Group;
  createImageStory: Story;
  createPost?: Maybe<Post>;
  createReel: Reel;
  createReelComment: ReelComment;
  createTextStory: Story;
  createUser: User;
  forgotPassword: Scalars['Boolean']['output'];
  inviteToGroup: Member;
  likePost?: Maybe<PostLike>;
  likeReel: ReelLike;
  likeReelComment: ReelCommentLike;
  likecomment?: Maybe<CommentLike>;
  rejectFriend?: Maybe<Friend>;
  resetPassword: User;
  sendMessage?: Maybe<Message>;
  sharePost?: Maybe<Scalars['String']['output']>;
  updateUser: User;
  updateUserBackground: User;
  updateUserProfile: User;
};


export type MutationAcceptFriendArgs = {
  friend: Scalars['ID']['input'];
};


export type MutationActivateUserArgs = {
  id: Scalars['String']['input'];
};


export type MutationAddFriendArgs = {
  friendInput: FriendInput;
};


export type MutationAuthenticateUserArgs = {
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
};


export type MutationCreateCommentArgs = {
  newComment: NewComment;
};


export type MutationCreateConversationArgs = {
  username: Scalars['String']['input'];
};


export type MutationCreateGroupArgs = {
  group: NewGroup;
};


export type MutationCreateImageStoryArgs = {
  input: NewImageStory;
};


export type MutationCreatePostArgs = {
  newPost: NewPost;
};


export type MutationCreateReelArgs = {
  reel: NewReel;
};


export type MutationCreateReelCommentArgs = {
  comment: NewReelComment;
};


export type MutationCreateTextStoryArgs = {
  input: NewTextStory;
};


export type MutationCreateUserArgs = {
  input: NewUser;
};


export type MutationForgotPasswordArgs = {
  email: Scalars['String']['input'];
};


export type MutationInviteToGroupArgs = {
  groupId: Scalars['ID']['input'];
  userId: Scalars['ID']['input'];
};


export type MutationLikePostArgs = {
  postID: Scalars['ID']['input'];
};


export type MutationLikeReelArgs = {
  reelId: Scalars['ID']['input'];
};


export type MutationLikeReelCommentArgs = {
  reelCommentId: Scalars['ID']['input'];
};


export type MutationLikecommentArgs = {
  commentID: Scalars['ID']['input'];
};


export type MutationRejectFriendArgs = {
  friend: Scalars['ID']['input'];
};


export type MutationResetPasswordArgs = {
  id: Scalars['String']['input'];
  password: Scalars['String']['input'];
};


export type MutationSendMessageArgs = {
  conversationID: Scalars['ID']['input'];
  image?: InputMaybe<Scalars['String']['input']>;
  message?: InputMaybe<Scalars['String']['input']>;
  postID?: InputMaybe<Scalars['ID']['input']>;
};


export type MutationSharePostArgs = {
  postID: Scalars['ID']['input'];
  userID: Scalars['ID']['input'];
};


export type MutationUpdateUserArgs = {
  input: UpdateUser;
};


export type MutationUpdateUserBackgroundArgs = {
  background: Scalars['String']['input'];
};


export type MutationUpdateUserProfileArgs = {
  profile: Scalars['String']['input'];
};

export type NewComment = {
  content: Scalars['String']['input'];
  parentComment?: InputMaybe<Scalars['ID']['input']>;
  parentPost?: InputMaybe<Scalars['ID']['input']>;
};

export type NewGroup = {
  about: Scalars['String']['input'];
  name: Scalars['String']['input'];
  privacy: Scalars['String']['input'];
};

export type NewImageStory = {
  image: Scalars['String']['input'];
};

export type NewPost = {
  content: Scalars['String']['input'];
  files?: InputMaybe<Array<InputMaybe<Scalars['String']['input']>>>;
  groupId?: InputMaybe<Scalars['String']['input']>;
  privacy: Scalars['String']['input'];
};

export type NewReel = {
  content: Scalars['String']['input'];
  video: Scalars['String']['input'];
};

export type NewReelComment = {
  content: Scalars['String']['input'];
  parentComment?: InputMaybe<Scalars['ID']['input']>;
  parentReel?: InputMaybe<Scalars['ID']['input']>;
};

export type NewTextStory = {
  color: Scalars['String']['input'];
  font: Scalars['String']['input'];
  text: Scalars['String']['input'];
};

export type NewUser = {
  dob: Scalars['Time']['input'];
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  gender: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type Pagination = {
  limit: Scalars['Int']['input'];
  start: Scalars['Int']['input'];
};

export type Post = {
  __typename?: 'Post';
  commentCount: Scalars['Int']['output'];
  comments?: Maybe<Array<Maybe<Comment>>>;
  content: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  files?: Maybe<Array<Maybe<Scalars['String']['output']>>>;
  groupId?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  likeCount: Scalars['Int']['output'];
  liked?: Maybe<Scalars['Boolean']['output']>;
  likes?: Maybe<Array<Maybe<PostLike>>>;
  privacy: Scalars['String']['output'];
  shareCount: Scalars['Int']['output'];
  user: User;
};

export type PostLike = {
  __typename?: 'PostLike';
  postId: Scalars['ID']['output'];
  user: User;
};

export type Query = {
  __typename?: 'Query';
  checkActivateLink: Scalars['Boolean']['output'];
  checkResetLink: Scalars['Boolean']['output'];
  getAuth: User;
  getCommentPost?: Maybe<Array<Maybe<Comment>>>;
  getConversations?: Maybe<Array<Maybe<Conversation>>>;
  getFriends?: Maybe<Array<Maybe<Friend>>>;
  getGroup?: Maybe<Group>;
  getGroupInvite?: Maybe<Array<Maybe<User>>>;
  getGroupPosts?: Maybe<Array<Maybe<Post>>>;
  getGroups?: Maybe<Array<Maybe<Group>>>;
  getJoinedGroups?: Maybe<Array<Maybe<Group>>>;
  getPost?: Maybe<Post>;
  getPosts?: Maybe<Array<Maybe<Post>>>;
  getReel: Reel;
  getReelComments: Array<Maybe<ReelComment>>;
  getReels?: Maybe<Array<Maybe<Scalars['ID']['output']>>>;
  getStories?: Maybe<Array<Maybe<Story>>>;
  getUser?: Maybe<User>;
  getUserFriends?: Maybe<Array<Maybe<Friend>>>;
  getUserMutuals?: Maybe<Array<Maybe<Friend>>>;
  getUserWithStories?: Maybe<Array<Maybe<User>>>;
  getUsers?: Maybe<Array<Maybe<User>>>;
};


export type QueryCheckActivateLinkArgs = {
  id: Scalars['String']['input'];
};


export type QueryCheckResetLinkArgs = {
  id: Scalars['String']['input'];
};


export type QueryGetCommentPostArgs = {
  postID: Scalars['ID']['input'];
};


export type QueryGetGroupArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetGroupInviteArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetGroupPostsArgs = {
  groupId: Scalars['ID']['input'];
  pagination: Pagination;
};


export type QueryGetPostArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetPostsArgs = {
  pagination: Pagination;
};


export type QueryGetReelArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGetReelCommentsArgs = {
  reelId: Scalars['ID']['input'];
};


export type QueryGetStoriesArgs = {
  username: Scalars['String']['input'];
};


export type QueryGetUserArgs = {
  username: Scalars['String']['input'];
};


export type QueryGetUserFriendsArgs = {
  username: Scalars['String']['input'];
};


export type QueryGetUserMutualsArgs = {
  username: Scalars['String']['input'];
};

export type Reel = {
  __typename?: 'Reel';
  commentCount: Scalars['Int']['output'];
  comments?: Maybe<Array<Maybe<ReelComment>>>;
  content: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  likeCount: Scalars['Int']['output'];
  liked?: Maybe<Scalars['Boolean']['output']>;
  likes?: Maybe<Array<Maybe<ReelLike>>>;
  shareCount: Scalars['Int']['output'];
  user: User;
  video: Scalars['String']['output'];
};

export type ReelComment = {
  __typename?: 'ReelComment';
  comments?: Maybe<Array<Maybe<ReelComment>>>;
  content: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  likeCount: Scalars['Int']['output'];
  liked?: Maybe<Scalars['Boolean']['output']>;
  likes?: Maybe<Array<Maybe<ReelCommentLike>>>;
  parentComment?: Maybe<ReelComment>;
  parentReel?: Maybe<Reel>;
  replyCount: Scalars['Int']['output'];
  user: User;
};

export type ReelCommentLike = {
  __typename?: 'ReelCommentLike';
  reelCommentId: Scalars['ID']['output'];
  user: User;
};

export type ReelLike = {
  __typename?: 'ReelLike';
  reelId: Scalars['ID']['output'];
  user: User;
};

export type Story = {
  __typename?: 'Story';
  color?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  font?: Maybe<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  image?: Maybe<Scalars['String']['output']>;
  text?: Maybe<Scalars['String']['output']>;
  user: User;
};

export type Subscription = {
  __typename?: 'Subscription';
  viewConversation?: Maybe<Array<Maybe<Message>>>;
};


export type SubscriptionViewConversationArgs = {
  conversationID: Scalars['ID']['input'];
};

export type UpdateUser = {
  firstName: Scalars['String']['input'];
  gender: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  password: Scalars['String']['input'];
};

export type User = {
  __typename?: 'User';
  active: Scalars['Boolean']['output'];
  background?: Maybe<Scalars['String']['output']>;
  createdAt: Scalars['Time']['output'];
  dob: Scalars['Time']['output'];
  email: Scalars['String']['output'];
  firstName: Scalars['String']['output'];
  friendCount: Scalars['Int']['output'];
  friended: Scalars['String']['output'];
  gender: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  lastName: Scalars['String']['output'];
  posts?: Maybe<Array<Maybe<Post>>>;
  profile?: Maybe<Scalars['String']['output']>;
  username: Scalars['String']['output'];
};

export type AcceptFriendMutationVariables = Exact<{
  friend: Scalars['ID']['input'];
}>;


export type AcceptFriendMutation = { __typename?: 'Mutation', acceptFriend?: { __typename?: 'Friend', accepted: boolean } | null };

export type AddFriendMutationVariables = Exact<{
  friendInput: FriendInput;
}>;


export type AddFriendMutation = { __typename?: 'Mutation', addFriend?: { __typename?: 'Friend', accepted: boolean, sender: { __typename?: 'User', username: string }, receiver: { __typename?: 'User', username: string } } | null };

export type GetFriendsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetFriendsQuery = { __typename?: 'Query', getFriends?: Array<{ __typename?: 'Friend', accepted: boolean, sender: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null }, receiver: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null } } | null> | null };

export type RejectFriendMutationVariables = Exact<{
  friend: Scalars['ID']['input'];
}>;


export type RejectFriendMutation = { __typename?: 'Mutation', rejectFriend?: { __typename?: 'Friend', accepted: boolean } | null };

export type CreateGroupMutationVariables = Exact<{
  group: NewGroup;
}>;


export type CreateGroupMutation = { __typename?: 'Mutation', createGroup: { __typename?: 'Group', id: string, name: string, about: string, privacy: string, background: string, members: Array<{ __typename?: 'Member', approved: boolean, role: string, user: { __typename?: 'User', firstName: string, lastName: string, username: string } }>, chat?: { __typename?: 'Conversation', id: string } | null } };

export type GetGroupQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetGroupQuery = { __typename?: 'Query', getGroup?: { __typename?: 'Group', id: string, name: string, about: string, privacy: string, background: string, isAdmin: boolean, memberCount: number, posts?: Array<{ __typename?: 'Post', id: string, content: string, privacy: string, likeCount: number, commentCount: number, shareCount: number, liked?: boolean | null, files?: Array<string | null> | null, createdAt: any, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null, email: string, gender: string, dob: any }, comments?: Array<{ __typename?: 'Comment', id: string, content: string } | null> | null } | null> | null, members: Array<{ __typename?: 'Member', approved: boolean, role: string, user: { __typename?: 'User', firstName: string, lastName: string, username: string } }>, chat?: { __typename?: 'Conversation', id: string } | null } | null };

export type GetGroupInviteQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetGroupInviteQuery = { __typename?: 'Query', getGroupInvite?: Array<{ __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null } | null> | null };

export type GetGroupPostsQueryVariables = Exact<{
  group: Scalars['ID']['input'];
  pagination: Pagination;
}>;


export type GetGroupPostsQuery = { __typename?: 'Query', getGroupPosts?: Array<{ __typename?: 'Post', id: string, content: string, privacy: string, likeCount: number, commentCount: number, shareCount: number, liked?: boolean | null, files?: Array<string | null> | null, createdAt: any, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null, email: string, gender: string, dob: any }, comments?: Array<{ __typename?: 'Comment', id: string, content: string } | null> | null } | null> | null };

export type GetGroupsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetGroupsQuery = { __typename?: 'Query', getGroups?: Array<{ __typename?: 'Group', id: string, name: string, about: string, privacy: string, background: string, memberCount: number, joined: string, members: Array<{ __typename?: 'Member', approved: boolean, role: string, user: { __typename?: 'User', firstName: string, lastName: string, username: string } }>, chat?: { __typename?: 'Conversation', id: string } | null } | null> | null };

export type GetJoinedGroupsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetJoinedGroupsQuery = { __typename?: 'Query', getJoinedGroups?: Array<{ __typename?: 'Group', id: string, name: string, about: string, privacy: string, background: string } | null> | null };

export type CreateConversationMutationVariables = Exact<{
  username: Scalars['String']['input'];
}>;


export type CreateConversationMutation = { __typename?: 'Mutation', createConversation?: { __typename?: 'Conversation', id: string } | null };

export type GetConversationsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetConversationsQuery = { __typename?: 'Query', getConversations?: Array<{ __typename?: 'Conversation', id: string, users: Array<{ __typename?: 'ConversationUsers', user: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null } }>, messages?: Array<{ __typename?: 'Message', message?: string | null } | null> | null } | null> | null };

export type SendMessageMutationVariables = Exact<{
  convID: Scalars['ID']['input'];
  message?: InputMaybe<Scalars['String']['input']>;
  image?: InputMaybe<Scalars['String']['input']>;
  post?: InputMaybe<Scalars['ID']['input']>;
}>;


export type SendMessageMutation = { __typename?: 'Mutation', sendMessage?: { __typename?: 'Message', id: string, message?: string | null, image?: string | null } | null };

export type ViewConversationSubscriptionVariables = Exact<{
  conversation: Scalars['ID']['input'];
}>;


export type ViewConversationSubscription = { __typename?: 'Subscription', viewConversation?: Array<{ __typename?: 'Message', message?: string | null, image?: string | null, createdAt: any, sender: { __typename?: 'User', firstName: string, lastName: string, username: string }, post?: { __typename?: 'Post', id: string, content: string, files?: Array<string | null> | null, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null } } | null } | null> | null };

export type CreateCommentMutationVariables = Exact<{
  newComment: NewComment;
}>;


export type CreateCommentMutation = { __typename?: 'Mutation', createComment?: { __typename?: 'Comment', id: string, content: string, liked?: boolean | null, likeCount: number, user: { __typename?: 'User', firstName: string, lastName: string, profile?: string | null } } | null };

export type CreatePostMutationVariables = Exact<{
  post: NewPost;
}>;


export type CreatePostMutation = { __typename?: 'Mutation', createPost?: { __typename?: 'Post', id: string, content: string, privacy: string, likeCount: number, commentCount: number, shareCount: number, liked?: boolean | null, files?: Array<string | null> | null, createdAt: any, user: { __typename?: 'User', firstName: string, lastName: string, profile?: string | null }, comments?: Array<{ __typename?: 'Comment', id: string, content: string } | null> | null } | null };

export type GetCommentPostQueryVariables = Exact<{
  postId: Scalars['ID']['input'];
}>;


export type GetCommentPostQuery = { __typename?: 'Query', getCommentPost?: Array<{ __typename?: 'Comment', id: string, content: string, liked?: boolean | null, likeCount: number, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null, email: string, gender: string, dob: any }, comments?: Array<{ __typename?: 'Comment', id: string, content: string, liked?: boolean | null, likeCount: number, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null, email: string, gender: string, dob: any } } | null> | null } | null> | null };

export type GetPostsQueryVariables = Exact<{
  pagination: Pagination;
}>;


export type GetPostsQuery = { __typename?: 'Query', getPosts?: Array<{ __typename?: 'Post', id: string, content: string, privacy: string, likeCount: number, commentCount: number, shareCount: number, liked?: boolean | null, files?: Array<string | null> | null, createdAt: any, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null, email: string, gender: string, dob: any }, comments?: Array<{ __typename?: 'Comment', id: string, content: string } | null> | null } | null> | null };

export type LikeCommentMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type LikeCommentMutation = { __typename?: 'Mutation', likecomment?: { __typename?: 'CommentLike', commentId: string } | null };

export type LikePostMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type LikePostMutation = { __typename?: 'Mutation', likePost?: { __typename?: 'PostLike', postId: string } | null };

export type SharePostMutationVariables = Exact<{
  user: Scalars['ID']['input'];
  post: Scalars['ID']['input'];
}>;


export type SharePostMutation = { __typename?: 'Mutation', sharePost?: string | null };

export type CreateReelMutationVariables = Exact<{
  reel: NewReel;
}>;


export type CreateReelMutation = { __typename?: 'Mutation', createReel: { __typename?: 'Reel', id: string, content: string, video: string, likeCount: number, user: { __typename?: 'User', firstName: string, lastName: string, username: string } } };

export type CreateReelCommentMutationVariables = Exact<{
  comment: NewReelComment;
}>;


export type CreateReelCommentMutation = { __typename?: 'Mutation', createReelComment: { __typename?: 'ReelComment', id: string, content: string, likeCount: number, replyCount: number, user: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null } } };

export type GetReelQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetReelQuery = { __typename?: 'Query', getReel: { __typename?: 'Reel', id: string, content: string, shareCount: number, likeCount: number, commentCount: number, liked?: boolean | null, video: string, user: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null } } };

export type GetReelCommentsQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetReelCommentsQuery = { __typename?: 'Query', getReelComments: Array<{ __typename?: 'ReelComment', id: string, content: string, likeCount: number, replyCount: number, liked?: boolean | null, user: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null }, comments?: Array<{ __typename?: 'ReelComment', id: string, content: string, likeCount: number, replyCount: number, liked?: boolean | null, user: { __typename?: 'User', firstName: string, lastName: string, username: string, profile?: string | null } } | null> | null } | null> };

export type GetReelsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetReelsQuery = { __typename?: 'Query', getReels?: Array<string | null> | null };

export type LikeReelMutationVariables = Exact<{
  reel: Scalars['ID']['input'];
}>;


export type LikeReelMutation = { __typename?: 'Mutation', likeReel: { __typename?: 'ReelLike', reelId: string } };

export type LikeReelCommentMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type LikeReelCommentMutation = { __typename?: 'Mutation', likeReelComment: { __typename?: 'ReelCommentLike', reelCommentId: string } };

export type CreateImageStoryMutationVariables = Exact<{
  story: NewImageStory;
}>;


export type CreateImageStoryMutation = { __typename?: 'Mutation', createImageStory: { __typename?: 'Story', id: string, text?: string | null, user: { __typename?: 'User', firstName: string, lastName: string, username: string } } };

export type CreateTextStoryMutationVariables = Exact<{
  story: NewTextStory;
}>;


export type CreateTextStoryMutation = { __typename?: 'Mutation', createTextStory: { __typename?: 'Story', id: string, text?: string | null, user: { __typename?: 'User', firstName: string, lastName: string, username: string } } };

export type GetStoriesQueryVariables = Exact<{
  username: Scalars['String']['input'];
}>;


export type GetStoriesQuery = { __typename?: 'Query', getStories?: Array<{ __typename?: 'Story', id: string, image?: string | null, text?: string | null, font?: string | null, color?: string | null } | null> | null };

export type GetUserWithStoriesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUserWithStoriesQuery = { __typename?: 'Query', getUserWithStories?: Array<{ __typename?: 'User', id: string, firstName: string, lastName: string, username: string, profile?: string | null } | null> | null };

export type ActivateUserMutationVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type ActivateUserMutation = { __typename?: 'Mutation', activateUser: { __typename?: 'User', id: string } };

export type AuthenticateUserMutationVariables = Exact<{
  email: Scalars['String']['input'];
  password: Scalars['String']['input'];
}>;


export type AuthenticateUserMutation = { __typename?: 'Mutation', authenticateUser: string };

export type CheckActivateLinkQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type CheckActivateLinkQuery = { __typename?: 'Query', checkActivateLink: boolean };

export type CheckResetLinkQueryVariables = Exact<{
  id: Scalars['String']['input'];
}>;


export type CheckResetLinkQuery = { __typename?: 'Query', checkResetLink: boolean };

export type CreateUserMutationVariables = Exact<{
  user: NewUser;
}>;


export type CreateUserMutation = { __typename?: 'Mutation', createUser: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, email: string, dob: any, gender: string, active: boolean } };

export type ForgotPasswordMutationVariables = Exact<{
  email: Scalars['String']['input'];
}>;


export type ForgotPasswordMutation = { __typename?: 'Mutation', forgotPassword: boolean };

export type GetUserQueryVariables = Exact<{
  username: Scalars['String']['input'];
}>;


export type GetUserQuery = { __typename?: 'Query', getUser?: { __typename?: 'User', id: string, firstName: string, lastName: string, username: string, email: string, dob: any, gender: string, active: boolean, profile?: string | null, background?: string | null, friended: string, friendCount: number, posts?: Array<{ __typename?: 'Post', id: string, content: string, privacy: string, likeCount: number, commentCount: number, shareCount: number, liked?: boolean | null, files?: Array<string | null> | null, createdAt: any, user: { __typename?: 'User', firstName: string, lastName: string, profile?: string | null } } | null> | null } | null };

export type ResetPasswordMutationVariables = Exact<{
  id: Scalars['String']['input'];
  password: Scalars['String']['input'];
}>;


export type ResetPasswordMutation = { __typename?: 'Mutation', resetPassword: { __typename?: 'User', id: string } };

export type UpdateUserMutationVariables = Exact<{
  updateUser: UpdateUser;
}>;


export type UpdateUserMutation = { __typename?: 'Mutation', updateUser: { __typename?: 'User', id: string } };

export type UpdateUserBackgroundMutationVariables = Exact<{
  background: Scalars['String']['input'];
}>;


export type UpdateUserBackgroundMutation = { __typename?: 'Mutation', updateUserBackground: { __typename?: 'User', id: string } };

export type UpdateUserProfileMutationVariables = Exact<{
  profile: Scalars['String']['input'];
}>;


export type UpdateUserProfileMutation = { __typename?: 'Mutation', updateUserProfile: { __typename?: 'User', id: string } };


export const AcceptFriendDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"acceptFriend"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"friend"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"acceptFriend"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"friend"},"value":{"kind":"Variable","name":{"kind":"Name","value":"friend"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accepted"}}]}}]}}]} as unknown as DocumentNode<AcceptFriendMutation, AcceptFriendMutationVariables>;
export const AddFriendDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addFriend"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"friendInput"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"FriendInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addFriend"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"friendInput"},"value":{"kind":"Variable","name":{"kind":"Name","value":"friendInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"accepted"}}]}}]}}]} as unknown as DocumentNode<AddFriendMutation, AddFriendMutationVariables>;
export const GetFriendsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getFriends"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getFriends"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"receiver"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"accepted"}}]}}]}}]} as unknown as DocumentNode<GetFriendsQuery, GetFriendsQueryVariables>;
export const RejectFriendDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"rejectFriend"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"friend"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"rejectFriend"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"friend"},"value":{"kind":"Variable","name":{"kind":"Name","value":"friend"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"accepted"}}]}}]}}]} as unknown as DocumentNode<RejectFriendMutation, RejectFriendMutationVariables>;
export const CreateGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"group"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewGroup"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"group"},"value":{"kind":"Variable","name":{"kind":"Name","value":"group"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"about"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"background"}},{"kind":"Field","name":{"kind":"Name","value":"members"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"approved"}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}},{"kind":"Field","name":{"kind":"Name","value":"chat"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<CreateGroupMutation, CreateGroupMutationVariables>;
export const GetGroupDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getGroup"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getGroup"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"about"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"background"}},{"kind":"Field","name":{"kind":"Name","value":"isAdmin"}},{"kind":"Field","name":{"kind":"Name","value":"posts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"commentCount"}},{"kind":"Field","name":{"kind":"Name","value":"shareCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"comments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}},{"kind":"Field","name":{"kind":"Name","value":"files"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}},{"kind":"Field","name":{"kind":"Name","value":"members"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"approved"}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}},{"kind":"Field","name":{"kind":"Name","value":"memberCount"}},{"kind":"Field","name":{"kind":"Name","value":"chat"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<GetGroupQuery, GetGroupQueryVariables>;
export const GetGroupInviteDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getGroupInvite"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getGroupInvite"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}}]}}]} as unknown as DocumentNode<GetGroupInviteQuery, GetGroupInviteQueryVariables>;
export const GetGroupPostsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getGroupPosts"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"group"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Pagination"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getGroupPosts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"groupId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"group"}}},{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"commentCount"}},{"kind":"Field","name":{"kind":"Name","value":"shareCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"comments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}},{"kind":"Field","name":{"kind":"Name","value":"files"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<GetGroupPostsQuery, GetGroupPostsQueryVariables>;
export const GetGroupsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"about"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"background"}},{"kind":"Field","name":{"kind":"Name","value":"members"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"approved"}},{"kind":"Field","name":{"kind":"Name","value":"role"}}]}},{"kind":"Field","name":{"kind":"Name","value":"memberCount"}},{"kind":"Field","name":{"kind":"Name","value":"joined"}},{"kind":"Field","name":{"kind":"Name","value":"chat"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]}}]} as unknown as DocumentNode<GetGroupsQuery, GetGroupsQueryVariables>;
export const GetJoinedGroupsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getJoinedGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getJoinedGroups"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"about"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"background"}}]}}]}}]} as unknown as DocumentNode<GetJoinedGroupsQuery, GetJoinedGroupsQueryVariables>;
export const CreateConversationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createConversation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"username"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createConversation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"username"},"value":{"kind":"Variable","name":{"kind":"Name","value":"username"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateConversationMutation, CreateConversationMutationVariables>;
export const GetConversationsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getConversations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getConversations"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"users"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"messages"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]}}]} as unknown as DocumentNode<GetConversationsQuery, GetConversationsQueryVariables>;
export const SendMessageDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"sendMessage"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"convID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"message"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"image"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"post"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sendMessage"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"conversationID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"convID"}}},{"kind":"Argument","name":{"kind":"Name","value":"message"},"value":{"kind":"Variable","name":{"kind":"Name","value":"message"}}},{"kind":"Argument","name":{"kind":"Name","value":"image"},"value":{"kind":"Variable","name":{"kind":"Name","value":"image"}}},{"kind":"Argument","name":{"kind":"Name","value":"postID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"post"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"image"}}]}}]}}]} as unknown as DocumentNode<SendMessageMutation, SendMessageMutationVariables>;
export const ViewConversationDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"subscription","name":{"kind":"Name","value":"viewConversation"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"conversation"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"viewConversation"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"conversationID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"conversation"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sender"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"message"}},{"kind":"Field","name":{"kind":"Name","value":"image"}},{"kind":"Field","name":{"kind":"Name","value":"post"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"files"}}]}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<ViewConversationSubscription, ViewConversationSubscriptionVariables>;
export const CreateCommentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createComment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"newComment"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewComment"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createComment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"newComment"},"value":{"kind":"Variable","name":{"kind":"Name","value":"newComment"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}}]}}]}}]} as unknown as DocumentNode<CreateCommentMutation, CreateCommentMutationVariables>;
export const CreatePostDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createPost"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"post"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewPost"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createPost"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"newPost"},"value":{"kind":"Variable","name":{"kind":"Name","value":"post"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"commentCount"}},{"kind":"Field","name":{"kind":"Name","value":"shareCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"comments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}},{"kind":"Field","name":{"kind":"Name","value":"files"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<CreatePostMutation, CreatePostMutationVariables>;
export const GetCommentPostDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getCommentPost"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"postId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getCommentPost"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"postID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"postId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"comments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}}]}}]}}]}}]}}]} as unknown as DocumentNode<GetCommentPostQuery, GetCommentPostQueryVariables>;
export const GetPostsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getPosts"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Pagination"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getPosts"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"commentCount"}},{"kind":"Field","name":{"kind":"Name","value":"shareCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"comments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"content"}}]}},{"kind":"Field","name":{"kind":"Name","value":"files"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<GetPostsQuery, GetPostsQueryVariables>;
export const LikeCommentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"likeComment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"likecomment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"commentID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"commentId"}}]}}]}}]} as unknown as DocumentNode<LikeCommentMutation, LikeCommentMutationVariables>;
export const LikePostDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"likePost"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"likePost"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"postID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"postId"}}]}}]}}]} as unknown as DocumentNode<LikePostMutation, LikePostMutationVariables>;
export const SharePostDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"sharePost"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"user"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"post"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"sharePost"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"user"}}},{"kind":"Argument","name":{"kind":"Name","value":"postID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"post"}}}]}]}}]} as unknown as DocumentNode<SharePostMutation, SharePostMutationVariables>;
export const CreateReelDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createReel"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"reel"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewReel"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createReel"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"reel"},"value":{"kind":"Variable","name":{"kind":"Name","value":"reel"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"video"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}}]}}]}}]} as unknown as DocumentNode<CreateReelMutation, CreateReelMutationVariables>;
export const CreateReelCommentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createReelComment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"comment"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewReelComment"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createReelComment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"comment"},"value":{"kind":"Variable","name":{"kind":"Name","value":"comment"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"replyCount"}}]}}]}}]} as unknown as DocumentNode<CreateReelCommentMutation, CreateReelCommentMutationVariables>;
export const GetReelDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getReel"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getReel"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"shareCount"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"commentCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"video"}}]}}]}}]} as unknown as DocumentNode<GetReelQuery, GetReelQueryVariables>;
export const GetReelCommentsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getReelComments"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getReelComments"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"reelId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"replyCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"comments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"replyCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}}]}}]}}]}}]} as unknown as DocumentNode<GetReelCommentsQuery, GetReelCommentsQueryVariables>;
export const GetReelsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getReels"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getReels"}}]}}]} as unknown as DocumentNode<GetReelsQuery, GetReelsQueryVariables>;
export const LikeReelDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"likeReel"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"reel"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"likeReel"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"reelId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"reel"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"reelId"}}]}}]}}]} as unknown as DocumentNode<LikeReelMutation, LikeReelMutationVariables>;
export const LikeReelCommentDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"likeReelComment"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"likeReelComment"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"reelCommentId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"reelCommentId"}}]}}]}}]} as unknown as DocumentNode<LikeReelCommentMutation, LikeReelCommentMutationVariables>;
export const CreateImageStoryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createImageStory"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"story"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewImageStory"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createImageStory"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"story"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"text"}}]}}]}}]} as unknown as DocumentNode<CreateImageStoryMutation, CreateImageStoryMutationVariables>;
export const CreateTextStoryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createTextStory"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"story"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewTextStory"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createTextStory"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"story"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}}]}},{"kind":"Field","name":{"kind":"Name","value":"text"}}]}}]}}]} as unknown as DocumentNode<CreateTextStoryMutation, CreateTextStoryMutationVariables>;
export const GetStoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getStories"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"username"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getStories"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"username"},"value":{"kind":"Variable","name":{"kind":"Name","value":"username"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"image"}},{"kind":"Field","name":{"kind":"Name","value":"text"}},{"kind":"Field","name":{"kind":"Name","value":"font"}},{"kind":"Field","name":{"kind":"Name","value":"color"}}]}}]}}]} as unknown as DocumentNode<GetStoriesQuery, GetStoriesQueryVariables>;
export const GetUserWithStoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetUserWithStories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUserWithStories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}}]}}]} as unknown as DocumentNode<GetUserWithStoriesQuery, GetUserWithStoriesQueryVariables>;
export const ActivateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"activateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"activateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ActivateUserMutation, ActivateUserMutationVariables>;
export const AuthenticateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"authenticateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"email"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"password"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"authenticateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"email"},"value":{"kind":"Variable","name":{"kind":"Name","value":"email"}}},{"kind":"Argument","name":{"kind":"Name","value":"password"},"value":{"kind":"Variable","name":{"kind":"Name","value":"password"}}}]}]}}]} as unknown as DocumentNode<AuthenticateUserMutation, AuthenticateUserMutationVariables>;
export const CheckActivateLinkDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"checkActivateLink"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"checkActivateLink"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}]}]}}]} as unknown as DocumentNode<CheckActivateLinkQuery, CheckActivateLinkQueryVariables>;
export const CheckResetLinkDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"checkResetLink"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"checkResetLink"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}]}]}}]} as unknown as DocumentNode<CheckResetLinkQuery, CheckResetLinkQueryVariables>;
export const CreateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"user"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"NewUser"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"user"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"active"}}]}}]}}]} as unknown as DocumentNode<CreateUserMutation, CreateUserMutationVariables>;
export const ForgotPasswordDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"forgotPassword"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"email"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"forgotPassword"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"email"},"value":{"kind":"Variable","name":{"kind":"Name","value":"email"}}}]}]}}]} as unknown as DocumentNode<ForgotPasswordMutation, ForgotPasswordMutationVariables>;
export const GetUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"username"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"getUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"username"},"value":{"kind":"Variable","name":{"kind":"Name","value":"username"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"dob"}},{"kind":"Field","name":{"kind":"Name","value":"gender"}},{"kind":"Field","name":{"kind":"Name","value":"active"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}},{"kind":"Field","name":{"kind":"Name","value":"background"}},{"kind":"Field","name":{"kind":"Name","value":"friended"}},{"kind":"Field","name":{"kind":"Name","value":"friendCount"}},{"kind":"Field","name":{"kind":"Name","value":"posts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"firstName"}},{"kind":"Field","name":{"kind":"Name","value":"lastName"}},{"kind":"Field","name":{"kind":"Name","value":"profile"}}]}},{"kind":"Field","name":{"kind":"Name","value":"content"}},{"kind":"Field","name":{"kind":"Name","value":"privacy"}},{"kind":"Field","name":{"kind":"Name","value":"likeCount"}},{"kind":"Field","name":{"kind":"Name","value":"commentCount"}},{"kind":"Field","name":{"kind":"Name","value":"shareCount"}},{"kind":"Field","name":{"kind":"Name","value":"liked"}},{"kind":"Field","name":{"kind":"Name","value":"files"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]}}]} as unknown as DocumentNode<GetUserQuery, GetUserQueryVariables>;
export const ResetPasswordDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"resetPassword"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"password"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"resetPassword"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"password"},"value":{"kind":"Variable","name":{"kind":"Name","value":"password"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ResetPasswordMutation, ResetPasswordMutationVariables>;
export const UpdateUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"updateUser"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UpdateUser"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"updateUser"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateUserMutation, UpdateUserMutationVariables>;
export const UpdateUserBackgroundDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateUserBackground"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"background"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUserBackground"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"background"},"value":{"kind":"Variable","name":{"kind":"Name","value":"background"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateUserBackgroundMutation, UpdateUserBackgroundMutationVariables>;
export const UpdateUserProfileDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"updateUserProfile"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"profile"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUserProfile"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"profile"},"value":{"kind":"Variable","name":{"kind":"Name","value":"profile"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<UpdateUserProfileMutation, UpdateUserProfileMutationVariables>;