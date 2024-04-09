export const typeDefs = `#graphql
type Game {
    id: ID!
    title: String
    platform: [String!]!
    reviews: [Review!]
}
type Review {
    id: ID!
    rating: Int!
    content: String!
    game: Game!
    author: Author!
}
type Author {
    id: ID!
    name: String!
    verified: Boolean!
    reviews: [Review!]
}
type Query {
    reviews: [Review]
    review(id: ID!): Review
    games: [Game]
    game(id:ID!):Game
    authors: [Author]
    author(id:ID!):Author
}
`

// user will only be able to start from Reviews only and then it can start jumping on other graphs from it if required/
// Query should must be there

// in platform array of fields can not be null and the fields can not be null too
// if we have used ! then that field is become required. it can not be null
// In graphql there are 5 scalar types that we an use
// int, float, string, boolean, ID
// ID field used by graphql for each data objects to identify.

// currently we are landing user on getting all the reviews, later on we will specify how can we land user to get only one review by id.

// Now we have added entry point to get the single review also.
