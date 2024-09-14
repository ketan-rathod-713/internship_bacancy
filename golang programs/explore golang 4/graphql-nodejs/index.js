import {ApolloServer} from "@apollo/server";
import {startStandaloneServer} from "@apollo/server/standalone"
import { typeDefs } from "./schema.js";
// server set up

// db
import db from "./_db.js"

// resolver functions for each properties we have defined in the typeDef of query.
const resolvers = {
    Query: {
        games(){
            return db.games
        },
        game(_, args){
            return db.games.find((game)=>{
                if(game.id == args.id){
                    return true
                }
            })
        },
        reviews(){
            return db.reviews
        },
         // we automatically get 3 arguments for such resolver functions
        // (parent, args, context)
        review(_,args){
            return db.reviews.find((review)=>{
                if(review.id == args.id){
                    return true
                }
            })
        },
        authors(){
            return db.authors
        },
        author(_,args){
            return db.authors.find((author)=>{
                if(author.id == args.id){
                    return true
                }
            })
        }
    },
    // How do we know which game we are reffering for this nested query
    // parent argument is the reference to the value returned by the parent resolver.
    // in our case parent will be a Game Object.
    Game: {
        reviews(parent) {
            return db.reviews.filter((review)=> {
                if(review.game_id === parent.id){
                    return true
                } else {
                    return false
                }
            })
        }
    }
}

const server = new ApolloServer({
    // typeDefs
    // resolvers
    typeDefs,
    resolvers
})

// user can make query like this 

/* 
game {
    title 
}

In this case we don't have to be worried about which fields are returned. Apollo server will handle it for us.
*/

const {url} = await startStandaloneServer(server, {
    port: 4000
})

console.log("Server ready at port", 4000)