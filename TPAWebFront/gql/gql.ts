/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n    mutation acceptFriend($friend: ID!) {\n        acceptFriend(friend: $friend) {\n            accepted\n        }\n    }\n": types.AcceptFriendDocument,
    "\n    mutation addFriend($friendInput: FriendInput!) {\n        addFriend(friendInput: $friendInput) {\n            sender {\n                username\n            }\n            receiver {\n                username\n            }\n            accepted\n        }\n    }\n": types.AddFriendDocument,
    "\n    query getFriends {\n        getFriends {\n            sender {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            receiver {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            accepted\n        }\n    }\n": types.GetFriendsDocument,
    "\n    mutation rejectFriend($friend: ID!) {\n        rejectFriend(friend: $friend) {\n            accepted\n        }\n    }\n": types.RejectFriendDocument,
    "\n    mutation createGroup($group: NewGroup!) {\n        createGroup(group: $group) {\n            id\n            name\n            about\n            privacy\n            background\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            chat {\n                id\n            }\n        }\n    }\n": types.CreateGroupDocument,
    "\n    query getGroup($id: ID!) {\n        getGroup(id: $id) {\n            id\n            name\n            about\n            privacy\n            background\n            posts {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                    email\n                    gender\n                    dob\n                }\n                content\n                privacy\n                likeCount\n                commentCount\n                shareCount\n                liked\n                comments {\n                    id\n                    content\n                }\n                files\n                createdAt\n            }\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            memberCount\n            chat {\n                id\n            }\n        }\n    }\n": types.GetGroupDocument,
    "\n    query getGroupPosts($group: ID!, $pagination: Pagination!) {\n        getGroupPosts(groupId: $group, pagination: $pagination) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n": types.GetGroupPostsDocument,
    "\n    query getGroups {\n        getGroups {\n            id\n            name\n            about\n            privacy\n            background\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            memberCount\n            joined\n            chat {\n                id\n            }\n        }\n    }\n": types.GetGroupsDocument,
    "\n    query getJoinedGroups {\n        getJoinedGroups {\n            id\n            name\n            about\n            privacy\n            background\n        }\n    }\n": types.GetJoinedGroupsDocument,
    "\n    mutation createConversation($username: String!) {\n        createConversation(username: $username) {\n            id\n        }\n    }\n": types.CreateConversationDocument,
    "\n    query getConversations {\n        getConversations {\n            id\n            users {\n                user {\n                    id\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n            }\n            messages {\n                message\n            }\n        }\n    }\n": types.GetConversationsDocument,
    "\n    mutation sendMessage($convID: ID!, $message: String, $image: String, $post: ID) {\n        sendMessage(conversationID: $convID, message: $message, image: $image, postID: $post) {\n            id\n            message\n            image\n        }\n    }\n": types.SendMessageDocument,
    "\n    subscription viewConversation($conversation: ID!) {\n        viewConversation(conversationID: $conversation) {\n            sender {\n                firstName\n                lastName\n                username\n            }\n            message\n            image\n            post {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n                content\n                files\n            }\n            createdAt\n        }\n    }\n": types.ViewConversationDocument,
    "\n    mutation createComment($newComment: NewComment!) {\n        createComment(newComment: $newComment) {\n            id\n            content\n            liked\n            likeCount\n            user {\n                firstName\n                lastName\n                profile\n            }\n        }\n    }\n": types.CreateCommentDocument,
    "\n    mutation createPost($post: NewPost!) {\n        createPost(newPost: $post) {\n            id\n            user {\n                firstName\n                lastName\n                profile\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n": types.CreatePostDocument,
    "\n    query getCommentPost($postId: ID!) {\n        getCommentPost(postID: $postId) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            liked\n            likeCount\n            comments {\n                id\n                content\n                liked\n                likeCount\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                    email\n                    gender\n                    dob\n                }\n            }\n        }\n    }\n": types.GetCommentPostDocument,
    "\n    query getPosts($pagination: Pagination!) {\n        getPosts(pagination: $pagination) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n": types.GetPostsDocument,
    "\n    mutation likeComment($id: ID!) {\n        likecomment(commentID: $id) {\n            commentId\n        }\n    }\n": types.LikeCommentDocument,
    "\n    mutation likePost($id: ID!) {\n        likePost(postID: $id) {\n            postId\n        }\n    }\n": types.LikePostDocument,
    "\n    mutation sharePost($user: ID!, $post: ID!) {\n        sharePost(userID: $user, postID: $post)\n    }\n": types.SharePostDocument,
    "\n    mutation createReel($reel: NewReel!) {\n        createReel(reel: $reel) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            content\n            video\n            likeCount\n        }\n    }\n": types.CreateReelDocument,
    "\n    mutation createReelComment($comment: NewReelComment!) {\n        createReelComment(comment: $comment) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            likeCount\n            replyCount\n        }\n    }\n": types.CreateReelCommentDocument,
    "\n    query getReel($id: ID!) {\n        getReel(id: $id) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            shareCount\n            likeCount\n            commentCount\n            liked\n            video\n        }\n    }\n": types.GetReelDocument,
    "\n    query getReelComments($id: ID!) {\n        getReelComments(reelId: $id) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            likeCount\n            replyCount\n            liked\n            comments {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n                content\n                likeCount\n                replyCount\n                liked\n            }\n        }\n    }\n": types.GetReelCommentsDocument,
    "\n    query getReels {\n        getReels\n    }\n": types.GetReelsDocument,
    "\n    mutation likeReel($reel: ID!) {\n        likeReel(reelId: $reel) {\n            reelId\n        }\n    }\n": types.LikeReelDocument,
    "\n    mutation likeReelComment($id: ID!) {\n        likeReelComment(reelCommentId: $id) {\n            reelCommentId\n        }\n    }\n": types.LikeReelCommentDocument,
    "\n    mutation createImageStory($story: NewImageStory!) {\n        createImageStory(input: $story) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            text\n        }\n    }\n": types.CreateImageStoryDocument,
    "\n    mutation createTextStory($story: NewTextStory!) {\n        createTextStory(input: $story) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            text\n        }\n    }\n": types.CreateTextStoryDocument,
    "\n    query getStories($username: String!) {\n        getStories(username: $username) {\n            id\n            image\n            text\n            font\n            color\n        }\n    }\n": types.GetStoriesDocument,
    "\n    query GetUserWithStories {\n        getUserWithStories {\n            id\n            firstName\n            lastName\n            username\n            profile\n        }\n    }\n": types.GetUserWithStoriesDocument,
    "\n    mutation activateUser($id: String!) {\n        activateUser(id: $id) {\n            id\n        }\n    }\n": types.ActivateUserDocument,
    "\n    mutation authenticateUser($email: String!, $password: String!) {\n        authenticateUser(email: $email, password: $password)\n    }\n": types.AuthenticateUserDocument,
    "\n    query checkActivateLink($id: String!) {\n        checkActivateLink(id: $id)\n    }\n": types.CheckActivateLinkDocument,
    "\n    query checkResetLink($id: String!) {\n        checkResetLink(id: $id)\n    }\n": types.CheckResetLinkDocument,
    "\n    mutation createUser($user: NewUser!) {\n        createUser(input: $user) {\n            id\n            firstName\n            lastName\n            username\n            email\n            dob\n            gender\n            active\n        }\n    }\n": types.CreateUserDocument,
    "\n    mutation forgotPassword($email: String!) {\n        forgotPassword(email: $email)\n    }\n": types.ForgotPasswordDocument,
    "\n    query getUser($username: String!) {\n        getUser(username: $username) {\n            id\n            firstName\n            lastName\n            username\n            email\n            dob\n            gender\n            active\n            profile\n            background\n            friended\n            friendCount\n            posts {\n                id\n                user {\n                    firstName\n                    lastName\n                    profile\n                }\n                content\n                privacy\n                likeCount\n                commentCount\n                shareCount\n                liked\n                files\n                createdAt\n            }\n        }\n    }\n": types.GetUserDocument,
    "\n    mutation resetPassword($id: String!, $password: String!) {\n        resetPassword(id: $id, password: $password) {\n            id\n        }\n    }\n": types.ResetPasswordDocument,
    "\n    mutation updateUser($updateUser: UpdateUser!) {\n        updateUser(input: $updateUser) {\n            id\n        }\n    }\n": types.UpdateUserDocument,
    "\n    mutation updateUserBackground($background: String!) {\n        updateUserBackground(background: $background) {\n            id\n        }\n    }\n": types.UpdateUserBackgroundDocument,
    "\n    mutation updateUserProfile($profile: String!) {\n        updateUserProfile(profile: $profile) {\n            id\n        }\n    }\n": types.UpdateUserProfileDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation acceptFriend($friend: ID!) {\n        acceptFriend(friend: $friend) {\n            accepted\n        }\n    }\n"): (typeof documents)["\n    mutation acceptFriend($friend: ID!) {\n        acceptFriend(friend: $friend) {\n            accepted\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation addFriend($friendInput: FriendInput!) {\n        addFriend(friendInput: $friendInput) {\n            sender {\n                username\n            }\n            receiver {\n                username\n            }\n            accepted\n        }\n    }\n"): (typeof documents)["\n    mutation addFriend($friendInput: FriendInput!) {\n        addFriend(friendInput: $friendInput) {\n            sender {\n                username\n            }\n            receiver {\n                username\n            }\n            accepted\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getFriends {\n        getFriends {\n            sender {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            receiver {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            accepted\n        }\n    }\n"): (typeof documents)["\n    query getFriends {\n        getFriends {\n            sender {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            receiver {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            accepted\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation rejectFriend($friend: ID!) {\n        rejectFriend(friend: $friend) {\n            accepted\n        }\n    }\n"): (typeof documents)["\n    mutation rejectFriend($friend: ID!) {\n        rejectFriend(friend: $friend) {\n            accepted\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createGroup($group: NewGroup!) {\n        createGroup(group: $group) {\n            id\n            name\n            about\n            privacy\n            background\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            chat {\n                id\n            }\n        }\n    }\n"): (typeof documents)["\n    mutation createGroup($group: NewGroup!) {\n        createGroup(group: $group) {\n            id\n            name\n            about\n            privacy\n            background\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            chat {\n                id\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getGroup($id: ID!) {\n        getGroup(id: $id) {\n            id\n            name\n            about\n            privacy\n            background\n            posts {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                    email\n                    gender\n                    dob\n                }\n                content\n                privacy\n                likeCount\n                commentCount\n                shareCount\n                liked\n                comments {\n                    id\n                    content\n                }\n                files\n                createdAt\n            }\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            memberCount\n            chat {\n                id\n            }\n        }\n    }\n"): (typeof documents)["\n    query getGroup($id: ID!) {\n        getGroup(id: $id) {\n            id\n            name\n            about\n            privacy\n            background\n            posts {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                    email\n                    gender\n                    dob\n                }\n                content\n                privacy\n                likeCount\n                commentCount\n                shareCount\n                liked\n                comments {\n                    id\n                    content\n                }\n                files\n                createdAt\n            }\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            memberCount\n            chat {\n                id\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getGroupPosts($group: ID!, $pagination: Pagination!) {\n        getGroupPosts(groupId: $group, pagination: $pagination) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n"): (typeof documents)["\n    query getGroupPosts($group: ID!, $pagination: Pagination!) {\n        getGroupPosts(groupId: $group, pagination: $pagination) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getGroups {\n        getGroups {\n            id\n            name\n            about\n            privacy\n            background\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            memberCount\n            joined\n            chat {\n                id\n            }\n        }\n    }\n"): (typeof documents)["\n    query getGroups {\n        getGroups {\n            id\n            name\n            about\n            privacy\n            background\n            members {\n                user {\n                    firstName\n                    lastName\n                    username\n                }\n                approved\n                role\n            }\n            memberCount\n            joined\n            chat {\n                id\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getJoinedGroups {\n        getJoinedGroups {\n            id\n            name\n            about\n            privacy\n            background\n        }\n    }\n"): (typeof documents)["\n    query getJoinedGroups {\n        getJoinedGroups {\n            id\n            name\n            about\n            privacy\n            background\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createConversation($username: String!) {\n        createConversation(username: $username) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation createConversation($username: String!) {\n        createConversation(username: $username) {\n            id\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getConversations {\n        getConversations {\n            id\n            users {\n                user {\n                    id\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n            }\n            messages {\n                message\n            }\n        }\n    }\n"): (typeof documents)["\n    query getConversations {\n        getConversations {\n            id\n            users {\n                user {\n                    id\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n            }\n            messages {\n                message\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation sendMessage($convID: ID!, $message: String, $image: String, $post: ID) {\n        sendMessage(conversationID: $convID, message: $message, image: $image, postID: $post) {\n            id\n            message\n            image\n        }\n    }\n"): (typeof documents)["\n    mutation sendMessage($convID: ID!, $message: String, $image: String, $post: ID) {\n        sendMessage(conversationID: $convID, message: $message, image: $image, postID: $post) {\n            id\n            message\n            image\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    subscription viewConversation($conversation: ID!) {\n        viewConversation(conversationID: $conversation) {\n            sender {\n                firstName\n                lastName\n                username\n            }\n            message\n            image\n            post {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n                content\n                files\n            }\n            createdAt\n        }\n    }\n"): (typeof documents)["\n    subscription viewConversation($conversation: ID!) {\n        viewConversation(conversationID: $conversation) {\n            sender {\n                firstName\n                lastName\n                username\n            }\n            message\n            image\n            post {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n                content\n                files\n            }\n            createdAt\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createComment($newComment: NewComment!) {\n        createComment(newComment: $newComment) {\n            id\n            content\n            liked\n            likeCount\n            user {\n                firstName\n                lastName\n                profile\n            }\n        }\n    }\n"): (typeof documents)["\n    mutation createComment($newComment: NewComment!) {\n        createComment(newComment: $newComment) {\n            id\n            content\n            liked\n            likeCount\n            user {\n                firstName\n                lastName\n                profile\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createPost($post: NewPost!) {\n        createPost(newPost: $post) {\n            id\n            user {\n                firstName\n                lastName\n                profile\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n"): (typeof documents)["\n    mutation createPost($post: NewPost!) {\n        createPost(newPost: $post) {\n            id\n            user {\n                firstName\n                lastName\n                profile\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getCommentPost($postId: ID!) {\n        getCommentPost(postID: $postId) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            liked\n            likeCount\n            comments {\n                id\n                content\n                liked\n                likeCount\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                    email\n                    gender\n                    dob\n                }\n            }\n        }\n    }\n"): (typeof documents)["\n    query getCommentPost($postId: ID!) {\n        getCommentPost(postID: $postId) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            liked\n            likeCount\n            comments {\n                id\n                content\n                liked\n                likeCount\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                    email\n                    gender\n                    dob\n                }\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getPosts($pagination: Pagination!) {\n        getPosts(pagination: $pagination) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n"): (typeof documents)["\n    query getPosts($pagination: Pagination!) {\n        getPosts(pagination: $pagination) {\n            id\n            user {\n                firstName\n                lastName\n                username\n                profile\n                email\n                gender\n                dob\n            }\n            content\n            privacy\n            likeCount\n            commentCount\n            shareCount\n            liked\n            comments {\n                id\n                content\n            }\n            files\n            createdAt\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation likeComment($id: ID!) {\n        likecomment(commentID: $id) {\n            commentId\n        }\n    }\n"): (typeof documents)["\n    mutation likeComment($id: ID!) {\n        likecomment(commentID: $id) {\n            commentId\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation likePost($id: ID!) {\n        likePost(postID: $id) {\n            postId\n        }\n    }\n"): (typeof documents)["\n    mutation likePost($id: ID!) {\n        likePost(postID: $id) {\n            postId\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation sharePost($user: ID!, $post: ID!) {\n        sharePost(userID: $user, postID: $post)\n    }\n"): (typeof documents)["\n    mutation sharePost($user: ID!, $post: ID!) {\n        sharePost(userID: $user, postID: $post)\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createReel($reel: NewReel!) {\n        createReel(reel: $reel) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            content\n            video\n            likeCount\n        }\n    }\n"): (typeof documents)["\n    mutation createReel($reel: NewReel!) {\n        createReel(reel: $reel) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            content\n            video\n            likeCount\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createReelComment($comment: NewReelComment!) {\n        createReelComment(comment: $comment) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            likeCount\n            replyCount\n        }\n    }\n"): (typeof documents)["\n    mutation createReelComment($comment: NewReelComment!) {\n        createReelComment(comment: $comment) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            likeCount\n            replyCount\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getReel($id: ID!) {\n        getReel(id: $id) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            shareCount\n            likeCount\n            commentCount\n            liked\n            video\n        }\n    }\n"): (typeof documents)["\n    query getReel($id: ID!) {\n        getReel(id: $id) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            shareCount\n            likeCount\n            commentCount\n            liked\n            video\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getReelComments($id: ID!) {\n        getReelComments(reelId: $id) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            likeCount\n            replyCount\n            liked\n            comments {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n                content\n                likeCount\n                replyCount\n                liked\n            }\n        }\n    }\n"): (typeof documents)["\n    query getReelComments($id: ID!) {\n        getReelComments(reelId: $id) {\n            id\n            user {\n                id\n                firstName\n                lastName\n                username\n                profile\n            }\n            content\n            likeCount\n            replyCount\n            liked\n            comments {\n                id\n                user {\n                    firstName\n                    lastName\n                    username\n                    profile\n                }\n                content\n                likeCount\n                replyCount\n                liked\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getReels {\n        getReels\n    }\n"): (typeof documents)["\n    query getReels {\n        getReels\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation likeReel($reel: ID!) {\n        likeReel(reelId: $reel) {\n            reelId\n        }\n    }\n"): (typeof documents)["\n    mutation likeReel($reel: ID!) {\n        likeReel(reelId: $reel) {\n            reelId\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation likeReelComment($id: ID!) {\n        likeReelComment(reelCommentId: $id) {\n            reelCommentId\n        }\n    }\n"): (typeof documents)["\n    mutation likeReelComment($id: ID!) {\n        likeReelComment(reelCommentId: $id) {\n            reelCommentId\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createImageStory($story: NewImageStory!) {\n        createImageStory(input: $story) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            text\n        }\n    }\n"): (typeof documents)["\n    mutation createImageStory($story: NewImageStory!) {\n        createImageStory(input: $story) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            text\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createTextStory($story: NewTextStory!) {\n        createTextStory(input: $story) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            text\n        }\n    }\n"): (typeof documents)["\n    mutation createTextStory($story: NewTextStory!) {\n        createTextStory(input: $story) {\n            id\n            user {\n                firstName\n                lastName\n                username\n            }\n            text\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getStories($username: String!) {\n        getStories(username: $username) {\n            id\n            image\n            text\n            font\n            color\n        }\n    }\n"): (typeof documents)["\n    query getStories($username: String!) {\n        getStories(username: $username) {\n            id\n            image\n            text\n            font\n            color\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query GetUserWithStories {\n        getUserWithStories {\n            id\n            firstName\n            lastName\n            username\n            profile\n        }\n    }\n"): (typeof documents)["\n    query GetUserWithStories {\n        getUserWithStories {\n            id\n            firstName\n            lastName\n            username\n            profile\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation activateUser($id: String!) {\n        activateUser(id: $id) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation activateUser($id: String!) {\n        activateUser(id: $id) {\n            id\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation authenticateUser($email: String!, $password: String!) {\n        authenticateUser(email: $email, password: $password)\n    }\n"): (typeof documents)["\n    mutation authenticateUser($email: String!, $password: String!) {\n        authenticateUser(email: $email, password: $password)\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query checkActivateLink($id: String!) {\n        checkActivateLink(id: $id)\n    }\n"): (typeof documents)["\n    query checkActivateLink($id: String!) {\n        checkActivateLink(id: $id)\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query checkResetLink($id: String!) {\n        checkResetLink(id: $id)\n    }\n"): (typeof documents)["\n    query checkResetLink($id: String!) {\n        checkResetLink(id: $id)\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation createUser($user: NewUser!) {\n        createUser(input: $user) {\n            id\n            firstName\n            lastName\n            username\n            email\n            dob\n            gender\n            active\n        }\n    }\n"): (typeof documents)["\n    mutation createUser($user: NewUser!) {\n        createUser(input: $user) {\n            id\n            firstName\n            lastName\n            username\n            email\n            dob\n            gender\n            active\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation forgotPassword($email: String!) {\n        forgotPassword(email: $email)\n    }\n"): (typeof documents)["\n    mutation forgotPassword($email: String!) {\n        forgotPassword(email: $email)\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    query getUser($username: String!) {\n        getUser(username: $username) {\n            id\n            firstName\n            lastName\n            username\n            email\n            dob\n            gender\n            active\n            profile\n            background\n            friended\n            friendCount\n            posts {\n                id\n                user {\n                    firstName\n                    lastName\n                    profile\n                }\n                content\n                privacy\n                likeCount\n                commentCount\n                shareCount\n                liked\n                files\n                createdAt\n            }\n        }\n    }\n"): (typeof documents)["\n    query getUser($username: String!) {\n        getUser(username: $username) {\n            id\n            firstName\n            lastName\n            username\n            email\n            dob\n            gender\n            active\n            profile\n            background\n            friended\n            friendCount\n            posts {\n                id\n                user {\n                    firstName\n                    lastName\n                    profile\n                }\n                content\n                privacy\n                likeCount\n                commentCount\n                shareCount\n                liked\n                files\n                createdAt\n            }\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation resetPassword($id: String!, $password: String!) {\n        resetPassword(id: $id, password: $password) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation resetPassword($id: String!, $password: String!) {\n        resetPassword(id: $id, password: $password) {\n            id\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation updateUser($updateUser: UpdateUser!) {\n        updateUser(input: $updateUser) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation updateUser($updateUser: UpdateUser!) {\n        updateUser(input: $updateUser) {\n            id\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation updateUserBackground($background: String!) {\n        updateUserBackground(background: $background) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation updateUserBackground($background: String!) {\n        updateUserBackground(background: $background) {\n            id\n        }\n    }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n    mutation updateUserProfile($profile: String!) {\n        updateUserProfile(profile: $profile) {\n            id\n        }\n    }\n"): (typeof documents)["\n    mutation updateUserProfile($profile: String!) {\n        updateUserProfile(profile: $profile) {\n            id\n        }\n    }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;