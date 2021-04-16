# KraickList
In this repository you will find simple web app for fictional startup called KraickList. This app will allow users to search ads from given sample data located in `data.gz`.

You could see the live version of this app [here](https://serene-depths-66275.herokuapp.com/). Try searching for "iPhone" to see some results.

## Changes
- **`Case insensitive`**, I changed the mechanism of this project to the insensitive case for query/keyword because the
  meaning of the search is to get all the "relevant" items according to the keyword and not limited to the case-sensitive.
- **`Add unit testing`**
- **`HTML Validation`**, I added simple HTML validation with an alert if the keyword is empty.
- **`Clean Architecture`**, I update the structure of this project from all functions in the main file to a separate directory according to the Clean Architecture structure.
- **`Misspelling Handle`**, I added misspelling feature to correct
the misspelled query. example search: ihone 8, hnda civic, tyota camry, etc.
- **`Ranked Result`**, Result gonna be ranked according to the 
most match with query.

## Deploy Method

I'm using Heroku to deploy this project.
