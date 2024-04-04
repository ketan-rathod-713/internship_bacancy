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
        reviews(){
            return db.reviews
        },
        authors(){
            return db.authors
        },

        // we automatically get 3 arguments for such resolver functions
        // (parent, args, context)
        review(_,args){
            return db.reviews.find((review)=>{
                if(review.id == args.id){
                    return true
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