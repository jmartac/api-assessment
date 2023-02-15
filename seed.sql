INSERT INTO users (created_at, updated_at, username, password_hash)
VALUES (NOW(), NOW(), 'existingUser',
        '$2a$10$U9qU/wflaZnEuGxE4skcZe136ERcOVSfe8d4mgevE.KDL9AalOHRW'), # password is 'megaP4ssword'
       (NOW(), NOW(), 'user2',
        '$2a$10$U9qU/wflaZnEuGxE4skcZe136ERcOVSfe8d4mgevE.KDL9AalOHRW'); # password is 'megaP4ssword'

INSERT INTO films (created_at, updated_at, title, director, release_date, genre, synopsis, user_id)
VALUES (NOW(), NOW(), 'The Shawshank Redemption', 'Frank Darabont', '1994-10-14', 'Drama',
        'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.',
        1),
       (NOW(), NOW(), 'The Godfather', 'Francis Ford Coppola', '1972-03-24', 'Crime, Drama',
        'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.',
        1),
       (NOW(), NOW(), 'The Dark Knight', 'Christopher Nolan', '2008-07-18', 'Action, Crime, Drama',
        'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, the caped crusader must come to terms with one of the greatest psychological tests of his ability to fight injustice.',
        1),
       (NOW(), NOW(), 'The Godfather: Part II', 'Francis Ford Coppola', '1974-12-20', 'Crime, Drama',
        'The early life and career of Vito Corleone in 1920s New York is portrayed while his son, Michael, expands and tightens his grip on his crime syndicate stretching from Lake Tahoe, Nevada to pre-revolution 1958 Cuba.',
        1),
       (NOW(), NOW(), 'Pulp Fiction', 'Quentin Tarantino', '1994-10-14', 'Crime, Drama',
        'The lives of two mob hitmen, a boxer, a gangster''s wife, and a pair of diner bandits intertwine in four tales of violence and redemption.',
        2),
       (NOW(), NOW(), 'The Lord of the Rings: The Return of the King', 'Peter Jackson', '2003-12-17',
        'Adventure, Drama, Fantasy',
        'Gandalf and Aragorn lead the World of Men against Sauron''s army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.',
        2),
       (NOW(), NOW(), 'The Good, the Bad and the Ugly', 'Sergio Leone', '1966-12-29', 'Western',
        'A bounty hunting scam joins two men in an uneasy alliance against a third in a race to find a fortune in gold buried in a remote cemetery.',
        2);
