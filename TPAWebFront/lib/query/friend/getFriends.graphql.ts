import { gql } from "@apollo/client";

export const GET_FRIENDS = gql`
    query getFriends($username: String!) {
        getFriends(username: $username) {
            id
            firstName
            lastName
            username
            email
            dob
            gender
            active
        }
    }
`;
