# Graphql 

## Introduction 

We query graphql server only on one route. totally different from the REST api.
for eg. mygraphql.com/graphql 

Our query can look like, 
```
Query {
    courses {
        id,
        title, 
        thumbnail_url
    }
}
```

This syntax allows us to get our required data from the server.


- What resource you want, just specify it.

```
query ReviewQuery {
    reviews {
        rating
    }
}
```

We have to define graphs in the server. which are connected. Hence we don't need to make a separate request to the server.

We can use any library for our graphql. here we are going to use appollo server. as it provides easy to get started with graphql server.

# Project 

```
npm init -y
npm pkg set type="module"
```

## Node Js Code

TypeDefs helps us to define the schema for our resources. Definitions of types of data.

Resolvers are the functions which we can pass it to appollo server which will do magic to it then.

Schema will clearly be same as that of database most probably.

Install graphql syntax highlighting. And then we can use it.

What if we want to get only single review or author. Here we can use variables.

First of all there should be entry point to get one review only.

Add another entry point to get single review entry point.

### Using Query Variables From The Frontend

```
query ReviewQuery($id:ID!){
  review(id: $id) {
    id,
    content
  }
}
```

Here we can pass all the variables at start that we want in our application and then call that particular method as it is as shown.

Now similarly define entry points for author and the game.

### Related Data In graphql

Every review has associated game and author.
Currently we haven't specified that relationship in our schema.

Now lets define this relationship inside our type definitions.

I can write a nested query from client likke this.

```
query GameQuery($id: ID!){
    game(id: $id){
        title,
        reviews {
            rating,
            content
        }
    }
}
```

But at the moment Appollo doesn't know how to handle this nested query thing. hence we have to specify it in resolver function.

Hence we arise a need for the resolver function for the nested types. But we can not put that in the Query because we don't want our nested resolvers to be entry point for the users.

So instead we create Game object inside resolvers and write a resolvers functions inside it.

```
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
```

# Cost of graphql and node js

They come with certain scaling and performance tradeoffs.

