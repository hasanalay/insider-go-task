Details are coming soon...

### Routes
````
	app.Get("/teams", handlers.GetTeams)
	app.Post("/teams", handlers.CreateTeam)
	app.Get("/teams/:id", handlers.GetTeamByID)
	app.Put("/teams/:id", handlers.UpdateTeam)
	app.Delete("/teams/:id", handlers.DeleteTeam)

	app.Get("/matches", handlers.GetMatches)
	app.Post("/matches", handlers.CreateMatch)
	app.Get("/matches/:id", handlers.GetMatchByID)
	app.Put("/matches/:id", handlers.UpdateMatch)
	app.Delete("/matches/:id", handlers.DeleteMatch)
	app.Get("matches/week/:week", handlers.GetMatchesByWeek)
````


### ENDPOINT Request Examples:

#### Create team
````
http://localhost:3000/teams

{
  "team_name": " team1 FC",
  "points": 0,
  "win": 0,
  "draw": 0,
  "lose": 0,
  "goal_difference": 0,
  "power": 9
}
````

#### Create match
````
http://localhost:3000/matches

{  "week": 1,
  "home_id": 3,
  "away_id": 4,
  "home_goals":0 ,
  "away_goals":0 ,
  "is_played": false 
  }
````

### Manually Create Fixture Data

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