# League API with GoLang and Postgresql

- This API is an implementation of Football league system. 
- The purpose of project is developing an API in a case study for  [insider](https://useinsider.com/)

The project is developed by [Hasan ALAY](https://github.com/hasanalay)

Project is deployed to the Google Cloud. You can acces from [here!](http://34.30.232.220:3000/api/matches). 

The project will be up in the cloud for 90 days. Due to free tier limitations. Here is public URL for API: http://34.30.232.220:3000

For a better experience a complete posman collection is uploaded in [repository](https://github.com/hasanalay/insider-go-task/blob/main/Insider-go-league-API.postman_collection.json).

### To run in your local machine 

1. Create an `.env` file by typing `touch .env` in your terminal while in project's root folder.
2. Enter these values. Dont forget the change your db connection information!
```
DB_HOST=yourHost
DB_PORT=yourPort
DB_USER=yourUser
DB_PASS=myStrongPass
DB_NAME=yourDB
DB_SSLMODE=disable
```
3. run `go mod tidy` in your terminal
4. run ` go run cmd/league-api/main.go` in your termianl to run project.
5. the application will accept requests in `http://localhost:3000/api/`

### Explore ENDPOINTS
```
  TEAMS
	Get("/teams", handlers.GetTeams)
	Post("/teams", handlers.CreateTeam)
	Get("/teams/:id", handlers.GetTeamByID)
	Put("/teams/:id", handlers.UpdateTeam)
	Delete("/teams/:id", handlers.DeleteTeam)

  MATCHES
	Get("/matches", handlers.GetMatches)
	Post("/matches", handlers.CreateMatch)
	Get("/matches/:id", handlers.GetMatchByID)
	Put("/matches/:id", handlers.UpdateMatch)
	Delete("/matches/:id", handlers.DeleteMatch)
	Get("/matches/week/:week", handlers.GetMatchesByWeek)

  LEAGUE
	Get("/league/:week", handlers.PlayMatch)
	Get("/league", handlers.PlayAllMatches)
	Put("/league/change-match/:id", handlers.ChangeMatchResult)
```

### Manually Create Data

Of course you can add data to tables by using create endpoints bu this is a faster way to test API.

Note: Make sure to add Teams first becuse teams and matches tables have relations!

#### Teams
```
INSERT INTO teams (team_name, points, win, draw, lose, goal_difference, power)
VALUES ('Team 1 FC', 0, 0, 0, 0, 0, 4);
VALUES ('Team 2 FC', 0, 0, 0, 0, 0, 5);
VALUES ('Team 3 FC', 0, 0, 0, 0, 0, 6);
VALUES ('Team 4 FC', 0, 0, 0, 0, 0, 3);
```
#### Matches
```
INSERT INTO matches (id, week, home_id, away_id, home_goals, away_goals, is_played)
VALUES
    (1, 1, 1, 2, 0, 0, false),
    (2, 1, 3, 4, 0, 0, false),
    
    (3, 2, 1, 3, 0, 0, false),
    (4, 2, 2, 4, 0, 0, false),
    
    (5, 3, 1, 4, 0, 0, false),
    (6, 3, 2, 3, 0, 0, false),
    
    (7, 4, 2, 1, 0, 0, false),
    (8, 4, 4, 3, 0, 0, false),
    
    (9 , 5, 3, 1, 0, 0, false),
    (10, 5, 4, 2, 0, 0, false),
    
    (11, 6, 4, 1, 0, 0, false),
    (12, 6, 3, 2, 0, 0, false);
```

### League ENDPOINTS

#### Play Matches of Week

`http://localhost:3000/api/league/4`

This endpoint is used for playing given week's matches. As Response it returns: 
- League leaderboard.
- Result of the matches.
- Next weeks's prediction of championship.( If week is greater than 4!)
```
{
    "league": {
        "league-leaderboard": [
            {
                "id": 3,
                "team_name": "Chelsea FC",
                "points": 8,
                "win": 2,
                "draw": 2,
                "lose": 0,
                "goal_difference": 3,
                "power": 5
            },
            {
                "id": 4,
                "team_name": "Liverpool FC",
                "points": 7,
                "win": 2,
                "draw": 1,
                "lose": 1,
                "goal_difference": 1,
                "power": 4
            },
            {
                "id": 2,
                "team_name": "Manchester United FC",
                "points": 6,
                "win": 2,
                "draw": 0,
                "lose": 2,
                "goal_difference": 3,
                "power": 7
            },
            {
                "id": 1,
                "team_name": "Manchester City FC",
                "points": 1,
                "win": 0,
                "draw": 1,
                "lose": 3,
                "goal_difference": -7,
                "power": 6
            }
        ]
    },
    "matches": {
        "matches": [
            {
                "id": 7,
                "week": 4,
                "home_id": 2,
                "away_id": 1,
                "home_goals": 8,
                "away_goals": 5,
                "is_played": true,
                "home_team": {
                    "id": 2,
                    "team_name": "Manchester United FC",
                    "points": 3,
                    "win": 1,
                    "draw": 0,
                    "lose": 2,
                    "goal_difference": 0,
                    "power": 7
                },
                "away_team": {
                    "id": 1,
                    "team_name": "Manchester City FC",
                    "points": 1,
                    "win": 0,
                    "draw": 1,
                    "lose": 2,
                    "goal_difference": -4,
                    "power": 6
                }
            },
            {
                "id": 8,
                "week": 4,
                "home_id": 4,
                "away_id": 3,
                "home_goals": 1,
                "away_goals": 1,
                "is_played": true,
                "home_team": {
                    "id": 4,
                    "team_name": "Liverpool FC",
                    "points": 6,
                    "win": 2,
                    "draw": 0,
                    "lose": 1,
                    "goal_difference": 1,
                    "power": 4
                },
                "away_team": {
                    "id": 3,
                    "team_name": "Chelsea FC",
                    "points": 7,
                    "win": 2,
                    "draw": 1,
                    "lose": 0,
                    "goal_difference": 3,
                    "power": 5
                }
            }
        ],
        "message": "Matches of week:4 played successfully!"
    },
    "next-week-predictions": {
        "message": "5th week predictions for Championship",
        "predictions": [
            {
                "Team": "Manchester United FC",
                "Percentage": "27%"
            },
            {
                "Team": "Manchester City FC",
                "Percentage": "4%"
            },
            {
                "Team": "Liverpool FC",
                "Percentage": "31%"
            },
            {
                "Team": "Chelsea FC",
                "Percentage": "36%"
            }
        ]
    }
}
```
#### Play All Matches

`http://localhost:3000/api/league`

This endpoint is used for playing all the matches at once. As Response it returns: 
- League leaderboard.
- Result of the all matches.
- Next weeks's prediction of championship.
```
{
    "league": {
        "league-leaderboard": [
            {
                "id": 2,
                "team_name": "Manchester United FC",
                "points": 10,
                "win": 3,
                "draw": 1,
                "lose": 2,
                "goal_difference": 6,
                "power": 7
            },
            {
                "id": 4,
                "team_name": "Liverpool FC",
                "points": 8,
                "win": 2,
                "draw": 2,
                "lose": 2,
                "goal_difference": -1,
                "power": 4
            },
            {
                "id": 3,
                "team_name": "Chelsea FC",
                "points": 8,
                "win": 2,
                "draw": 2,
                "lose": 2,
                "goal_difference": -3,
                "power": 5
            },
            {
                "id": 1,
                "team_name": "Manchester City FC",
                "points": 7,
                "win": 2,
                "draw": 1,
                "lose": 3,
                "goal_difference": -2,
                "power": 6
            }
        ]
    },
    "matches": {
        "matches": [
            {
                "id": 2,
                "week": 1,
                "home_id": 4,
                "away_id": 3,
                "home_goals": 3,
                "away_goals": 4,
                "is_played": true,
                "home_team": {
                    "id": 4,
                    "team_name": "Liverpool FC",
                    "points": 7,
                    "win": 2,
                    "draw": 1,
                    "lose": 1,
                    "goal_difference": 1,
                    "power": 4
                },
                "away_team": {
                    "id": 3,
                    "team_name": "Chelsea FC",
                    "points": 8,
                    "win": 2,
                    "draw": 2,
                    "lose": 0,
                    "goal_difference": 3,
                    "power": 5
                }
            },
            {},
            {},
            {},
            {},
            {},
            {},
            {},
            {},
            {},
            {},
            {}
        ],
        "message": "All Matches played successfully!"
    },
    "next-week-predictions": {
        "message": "Next week match predictions",
        "predictions": [
            {
                "Team": "Liverpool FC",
                "Percentage": "24%"
            },
            {
                "Team": "Manchester City FC",
                "Percentage": "21%"
            },
            {
                "Team": "Chelsea FC",
                "Percentage": "24%"
            },
            {
                "Team": "Manchester United FC",
                "Percentage": "30%"
            }
        ]
    }
}
```

#### Change Match Scores

`http://localhost:3000/api/league/change-match/1`

This endpoint is used for playing all the matches at once. As request it takes a match object like this:
```
{  
  "week": 1,
  "home_id": 1,
  "away_id": 2,
  "home_goals":0,
  "away_goals":3,
  "is_played": true 
}
```
As Response it returns: 
- League leaderboard.
- Result of the all matches.
- Next weeks's prediction of championship.
```
{
    "league": {
        "league-leaderboard": [
            {
                "id": 2,
                "team_name": "Manchester United FC",
                "points": 10,
                "win": 3,
                "draw": 1,
                "lose": 2,
                "goal_difference": 3,
                "power": 7
            },
            {
                "id": 4,
                "team_name": "Liverpool FC",
                "points": 8,
                "win": 2,
                "draw": 2,
                "lose": 2,
                "goal_difference": -1,
                "power": 4
            },
            {
                "id": 3,
                "team_name": "Chelsea FC",
                "points": 8,
                "win": 2,
                "draw": 2,
                "lose": 2,
                "goal_difference": -3,
                "power": 5
            },
            {
                "id": 1,
                "team_name": "Manchester City FC",
                "points": 7,
                "win": 2,
                "draw": 1,
                "lose": 3,
                "goal_difference": -3,
                "power": 6
            }
        ]
    },
    "match": {
        "id": 1,
        "week": 1,
        "home_id": 1,
        "away_id": 2,
        "home_goals": 0,
        "away_goals": 3,
        "is_played": true,
        "home_team": {
            "id": 1,
            "team_name": "Manchester City FC",
            "points": 7,
            "win": 2,
            "draw": 1,
            "lose": 3,
            "goal_difference": -3,
            "power": 6
        },
        "away_team": {
            "id": 2,
            "team_name": "Manchester United FC",
            "points": 10,
            "win": 3,
            "draw": 1,
            "lose": 2,
            "goal_difference": 3,
            "power": 7
        }
    },
    "message": "Match updated successfully!"
}
```


