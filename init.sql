-- Enable PostGIS extension for geographic queries
CREATE EXTENSION IF NOT EXISTS postgis;

-- Connect to the database
\c theatre_api;

-- Create tables (matching GORM models)
CREATE TABLE IF NOT EXISTS locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    latitude DECIMAL(10,8) NOT NULL,
    longitude DECIMAL(11,8) NOT NULL,
    postal_code VARCHAR(20),
    address TEXT,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS theatre_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS show_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS theatres (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    capacity INTEGER NOT NULL CHECK (capacity > 0),
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    website VARCHAR(255),
    image_url TEXT,
    is_featured BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    location_id UUID NOT NULL REFERENCES locations(id),
    theatre_type_id UUID NOT NULL REFERENCES theatre_types(id)
);

CREATE TABLE IF NOT EXISTS shows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    director VARCHAR(255),
    "cast" TEXT,
    duration INTEGER NOT NULL CHECK (duration > 0),
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    price DECIMAL(10,2) CHECK (price >= 0),
    image_url TEXT,
    trailer_url TEXT,
    is_featured BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    theatre_id UUID NOT NULL REFERENCES theatres(id),
    show_type_id UUID NOT NULL REFERENCES show_types(id),
    CHECK (end_date IS NULL OR start_date IS NULL OR end_date >= start_date)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_locations_coordinates ON locations(latitude, longitude);
CREATE INDEX IF NOT EXISTS idx_locations_active ON locations(is_active);
CREATE INDEX IF NOT EXISTS idx_theatres_location ON theatres(location_id);
CREATE INDEX IF NOT EXISTS idx_theatres_type ON theatres(theatre_type_id);
CREATE INDEX IF NOT EXISTS idx_theatres_active ON theatres(is_active);
CREATE INDEX IF NOT EXISTS idx_theatres_featured ON theatres(is_featured);
CREATE INDEX IF NOT EXISTS idx_shows_theatre ON shows(theatre_id);
CREATE INDEX IF NOT EXISTS idx_shows_type ON shows(show_type_id);
CREATE INDEX IF NOT EXISTS idx_shows_active ON shows(is_active);
CREATE INDEX IF NOT EXISTS idx_shows_featured ON shows(is_featured);
CREATE INDEX IF NOT EXISTS idx_shows_dates ON shows(start_date, end_date);

-- Insert sample locations (expanded to support 100+ theatres)
INSERT INTO locations (id, name, city, state, country, latitude, longitude, postal_code, address, description, is_active, created_at, updated_at) VALUES
(gen_random_uuid(), 'Manhattan Theater District', 'New York', 'New York', 'United States', 40.7589, -73.9851, '10036', 'Times Square, NYC', 'Heart of Broadway theater district', true, NOW(), NOW()),
(gen_random_uuid(), 'Upper West Side', 'New York', 'New York', 'United States', 40.7831, -73.9712, '10019', 'Upper West Side, NYC', 'Cultural hub with Lincoln Center', true, NOW(), NOW()),
(gen_random_uuid(), 'Lower Manhattan', 'New York', 'New York', 'United States', 40.7074, -74.0113, '10013', 'Lower Manhattan, NYC', 'Historic downtown theater scene', true, NOW(), NOW()),
(gen_random_uuid(), 'West End', 'London', 'England', 'United Kingdom', 51.5142, -0.1506, 'WC2E', 'West End, London', 'London''s premier theater district', true, NOW(), NOW()),
(gen_random_uuid(), 'South Bank', 'London', 'England', 'United Kingdom', 51.5074, -0.0977, 'SE1', 'South Bank, London', 'Modern theater complex on Thames', true, NOW(), NOW()),
(gen_random_uuid(), 'Chicago Loop', 'Chicago', 'Illinois', 'United States', 41.8781, -87.6298, '60601', 'Loop, Chicago, IL', 'Historic theater district in downtown Chicago', true, NOW(), NOW()),
(gen_random_uuid(), 'North Side Chicago', 'Chicago', 'Illinois', 'United States', 41.9484, -87.6553, '60614', 'North Side, Chicago, IL', 'Vibrant neighborhood theater scene', true, NOW(), NOW()),
(gen_random_uuid(), 'Las Vegas Strip', 'Las Vegas', 'Nevada', 'United States', 36.1162, -115.1739, '89109', 'Las Vegas Strip, NV', 'Entertainment capital with world-class shows', true, NOW(), NOW()),
(gen_random_uuid(), 'Downtown Las Vegas', 'Las Vegas', 'Nevada', 'United States', 36.1699, -115.1398, '89101', 'Downtown Las Vegas, NV', 'Historic entertainment district', true, NOW(), NOW()),
(gen_random_uuid(), 'Toronto Entertainment District', 'Toronto', 'Ontario', 'Canada', 43.6426, -79.3871, 'M5V', 'King Street West, Toronto', 'Canada''s largest theater district', true, NOW(), NOW()),
(gen_random_uuid(), 'Hollywood', 'Los Angeles', 'California', 'United States', 34.0928, -118.3287, '90028', 'Hollywood, CA', 'Entertainment capital of the world', true, NOW(), NOW()),
(gen_random_uuid(), 'Beverly Hills', 'Los Angeles', 'California', 'United States', 34.0736, -118.4004, '90210', 'Beverly Hills, CA', 'Upscale theater venues', true, NOW(), NOW()),
(gen_random_uuid(), 'Boston Theater District', 'Boston', 'Massachusetts', 'United States', 42.3505, -71.0621, '02116', 'Theater District, Boston, MA', 'Historic New England theater hub', true, NOW(), NOW()),
(gen_random_uuid(), 'Philadelphia Center City', 'Philadelphia', 'Pennsylvania', 'United States', 39.9526, -75.1652, '19102', 'Center City, Philadelphia, PA', 'Cultural center of Philadelphia', true, NOW(), NOW()),
(gen_random_uuid(), 'San Francisco Union Square', 'San Francisco', 'California', 'United States', 37.7877, -122.4074, '94108', 'Union Square, San Francisco, CA', 'Premier theater district', true, NOW(), NOW()),
(gen_random_uuid(), 'Washington DC Kennedy Center', 'Washington', 'District of Columbia', 'United States', 38.8955, -77.0563, '20566', 'Kennedy Center, Washington DC', 'National cultural center', true, NOW(), NOW()),
(gen_random_uuid(), 'Atlanta Midtown', 'Atlanta', 'Georgia', 'United States', 33.7839, -84.3830, '30309', 'Midtown Atlanta, GA', 'Cultural district of the South', true, NOW(), NOW()),
(gen_random_uuid(), 'Seattle Capitol Hill', 'Seattle', 'Washington', 'United States', 47.6205, -122.3212, '98102', 'Capitol Hill, Seattle, WA', 'Artistic neighborhood with theaters', true, NOW(), NOW()),
(gen_random_uuid(), 'Denver Arts District', 'Denver', 'Colorado', 'United States', 39.7392, -104.9903, '80202', 'Arts District, Denver, CO', 'Growing theater scene in the Rockies', true, NOW(), NOW()),
(gen_random_uuid(), 'Miami Beach', 'Miami', 'Florida', 'United States', 25.7907, -80.1300, '33139', 'Miami Beach, FL', 'Tropical theater destination', true, NOW(), NOW());

-- Insert sample theatre types
INSERT INTO theatre_types (id, name, description, is_active, created_at, updated_at) VALUES
(gen_random_uuid(), 'Broadway', 'Professional theaters in Manhattan''s Theater District with 500+ seats', true, NOW(), NOW()),
(gen_random_uuid(), 'Off-Broadway', 'Professional theaters in NYC with 100-499 seats', true, NOW(), NOW()),
(gen_random_uuid(), 'Off-Off-Broadway', 'Intimate theaters in NYC with fewer than 100 seats', true, NOW(), NOW()),
(gen_random_uuid(), 'Regional Theater', 'Professional theaters outside of NYC', true, NOW(), NOW()),
(gen_random_uuid(), 'Community Theater', 'Non-professional theaters run by local communities', true, NOW(), NOW()),
(gen_random_uuid(), 'Dinner Theater', 'Theaters that serve meals during performances', true, NOW(), NOW()),
(gen_random_uuid(), 'Arena Theater', 'Theater-in-the-round with audience surrounding the stage', true, NOW(), NOW()),
(gen_random_uuid(), 'Outdoor Theater', 'Open-air venues for summer performances', true, NOW(), NOW()),
(gen_random_uuid(), 'Concert Hall', 'Large venues primarily for musical performances', true, NOW(), NOW()),
(gen_random_uuid(), 'Opera House', 'Venues specifically designed for opera performances', true, NOW(), NOW());

-- Insert sample show types
INSERT INTO show_types (id, name, description, is_active, created_at, updated_at) VALUES
(gen_random_uuid(), 'Musical', 'Theatrical productions featuring songs, spoken dialogue, acting, and dance', true, NOW(), NOW()),
(gen_random_uuid(), 'Play', 'Dramatic works performed by actors on stage', true, NOW(), NOW()),
(gen_random_uuid(), 'Opera', 'Musical theater combining singing, orchestral music, and dramatic performance', true, NOW(), NOW()),
(gen_random_uuid(), 'Ballet', 'Classical dance performances with orchestral accompaniment', true, NOW(), NOW()),
(gen_random_uuid(), 'Comedy', 'Humorous performances designed to entertain and amuse', true, NOW(), NOW()),
(gen_random_uuid(), 'Drama', 'Serious theatrical works exploring complex themes and emotions', true, NOW(), NOW()),
(gen_random_uuid(), 'Concert', 'Musical performances by orchestras, bands, or solo artists', true, NOW(), NOW()),
(gen_random_uuid(), 'Variety Show', 'Entertainment featuring multiple acts and performers', true, NOW(), NOW()),
(gen_random_uuid(), 'Cabaret', 'Intimate performances combining music, dance, and comedy', true, NOW(), NOW()),
(gen_random_uuid(), 'Children''s Theater', 'Performances specifically designed for young audiences', true, NOW(), NOW()),
(gen_random_uuid(), 'Experimental', 'Avant-garde and innovative theatrical works', true, NOW(), NOW()),
(gen_random_uuid(), 'Dance Performance', 'Contemporary and modern dance shows', true, NOW(), NOW());

-- Create a function to insert theaters and shows
DO $$
DECLARE
    location_ids uuid[];
    theatre_type_ids uuid[];
    show_type_ids uuid[];
    current_theatre_id uuid;
    i integer;
    j integer;
    theatre_names text[] := ARRAY[
        'Majestic Theatre', 'Palace Theatre', 'Lyceum Theatre', 'Broadhurst Theatre', 'Imperial Theatre',
        'Ambassador Theatre', 'Eugene O''Neill Theatre', 'Gershwin Theatre', 'Minskoff Theatre', 'New Amsterdam Theatre',
        'Richard Rodgers Theatre', 'St. James Theatre', 'Winter Garden Theatre', 'Nederlander Theatre', 'Brooks Atkinson Theatre',
        'Bernard B. Jacobs Theatre', 'Gerald Schoenfeld Theatre', 'Booth Theatre', 'Music Box Theatre', 'Belasco Theatre',
        'Longacre Theatre', 'Lunt-Fontanne Theatre', 'Al Hirschfeld Theatre', 'August Wilson Theatre', 'Barrymore Theatre',
        'Biltmore Theatre', 'Circle in the Square Theatre', 'Cort Theatre', 'Ethel Barrymore Theatre', 'Hayes Theater',
        'Helen Hayes Theatre', 'Hudson Theatre', 'John Golden Theatre', 'Marquis Theatre', 'Neil Simon Theatre',
        'Samuel J. Friedman Theatre', 'Stephen Sondheim Theatre', 'Studio 54', 'TKTS Red Steps', 'Vivian Beaumont Theater',
        'Apollo Theatre', 'Dominion Theatre', 'Her Majesty''s Theatre', 'London Palladium', 'Lyceum Theatre London',
        'National Theatre', 'Old Vic Theatre', 'Phoenix Theatre', 'Prince Edward Theatre', 'Queen''s Theatre',
        'Savoy Theatre', 'Theatre Royal Drury Lane', 'Victoria Palace Theatre', 'Wyndham''s Theatre', 'Adelphi Theatre',
        'Chicago Theatre', 'Oriental Theatre', 'Cadillac Palace Theatre', 'CIBC Theatre', 'Goodman Theatre',
        'Steppenwolf Theatre', 'Second City', 'Victory Gardens Theater', 'Lookingglass Theatre', 'Chicago Shakespeare Theater',
        'Bellagio Theater', 'Colosseum at Caesars Palace', 'MGM Grand Garden Arena', 'Park Theater', 'Zappos Theater',
        'Smith Center', 'Le Reve Theater', 'Blue Man Group Theater', 'Jubilee Theater', 'Penn & Teller Theater',
        'Princess of Wales Theatre', 'Royal Alexandra Theatre', 'Ed Mirvish Theatre', 'Elgin Theatre', 'Winter Garden Theatre Toronto',
        'Panasonic Theatre', 'CAA Theatre', 'Danforth Music Hall', 'Phoenix Concert Theatre', 'The Opera House',
        'Dolby Theatre', 'TCL Chinese Theatre', 'El Capitan Theatre', 'Pantages Theatre', 'Ahmanson Theatre',
        'Mark Taper Forum', 'Geffen Playhouse', 'Kirk Douglas Theatre', 'Wallis Annenberg Center', 'Hollywood Bowl',
        'Boston Opera House', 'Wang Theatre', 'Berklee Performance Center', 'Huntington Theatre', 'American Repertory Theater',
        'Kimmel Center', 'Academy of Music', 'Walnut Street Theatre', 'Forrest Theatre', 'Merriam Theater',
        'Curran Theatre', 'Golden Gate Theatre', 'Orpheum Theatre SF', 'American Conservatory Theater', 'Magic Theatre',
        'Kennedy Center Opera House', 'Kennedy Center Concert Hall', 'Arena Stage', 'Shakespeare Theatre Company', 'Studio Theatre',
        'Fox Theatre Atlanta', 'Alliance Theatre', 'Actor''s Express', 'Horizon Theatre', 'Dad''s Garage Theatre',
        'Paramount Theatre Seattle', 'ACT Theatre', 'Cornish Playhouse', 'Moore Theatre', 'Neptune Theatre',
        'Denver Center for Performing Arts', 'Buell Theatre', 'Ellie Caulkins Opera House', 'Space Theatre', 'Stage Theatre'
    ];
    show_titles text[] := ARRAY[
        'The Lion King', 'Hamilton', 'The Phantom of the Opera', 'Chicago', 'Cats', 'Les Misérables', 'Wicked',
        'The Book of Mormon', 'Dear Evan Hansen', 'Come From Away', 'Hadestown', 'Frozen', 'Aladdin', 'The King and I',
        'West Side Story', 'My Fair Lady', 'Oklahoma!', 'South Pacific', 'The Sound of Music', 'Carousel',
        'Romeo and Juliet', 'Hamlet', 'Macbeth', 'King Lear', 'Othello', 'A Midsummer Night''s Dream', 'The Tempest',
        'Death of a Salesman', 'A Streetcar Named Desire', 'The Glass Menagerie', 'Cat on a Hot Tin Roof', 'Long Day''s Journey Into Night',
        'Our Town', 'The Crucible', 'Angels in America', 'Rent', 'Avenue Q', 'Spring Awakening', 'Next to Normal',
        'In the Heights', 'Memphis', 'The Color Purple', 'Waitress', 'Beautiful', 'Carole King Musical', 'Jersey Boys',
        'Mamma Mia!', 'ABBA Musical', 'We Will Rock You', 'Queen Musical', 'Billy Elliot', 'Matilda', 'Charlie and the Chocolate Factory',
        'Mary Poppins', 'The Little Mermaid', 'Beauty and the Beast', 'Tarzan', 'The Jungle Book', 'Peter Pan',
        'A Christmas Carol', 'The Nutcracker', 'Swan Lake', 'Giselle', 'Romeo and Juliet Ballet', 'Sleeping Beauty',
        'La Bohème', 'Carmen', 'The Magic Flute', 'Don Giovanni', 'Tosca', 'Madama Butterfly', 'La Traviata',
        'The Barber of Seville', 'Rigoletto', 'Aida', 'Turandot', 'The Marriage of Figaro', 'Così fan tutte',
        'Blue Man Group', 'Cirque du Soleil: O', 'Penn & Teller', 'David Copperfield', 'Absinthe', 'Le Rêve',
        'Thunder from Down Under', 'Chippendales', 'Magic Mike Live', 'Zombie Burlesque', 'Menopause the Musical',
        'Tony n'' Tina''s Wedding', 'Late Night Catechism', 'Defending the Caveman', 'I Love You, You''re Perfect, Now Change',
        'Stomp', 'Tap Dogs', 'Lord of the Dance', 'Riverdance', 'Celtic Thunder', 'Celtic Woman',
        'The Second City', 'Saturday Night Live', 'Comedy Central Presents', 'Whose Line Is It Anyway?', 'Improv Comedy',
        'Stand-Up Comedy Night', 'Open Mic Night', 'Comedy Roast', 'Sketch Comedy Show', 'Musical Comedy Revue'
    ];
    directors text[] := ARRAY[
        'Julie Taymor', 'Thomas Kail', 'Hal Prince', 'Trevor Nunn', 'Cameron Mackintosh', 'Andrew Lloyd Webber',
        'Stephen Sondheim', 'Lin-Manuel Miranda', 'Jonathan Larson', 'Jason Robert Brown', 'Pasek and Paul',
        'Alan Menken', 'Tim Rice', 'Elton John', 'ABBA', 'Queen', 'Billy Joel', 'Carole King',
        'William Shakespeare', 'Tennessee Williams', 'Arthur Miller', 'Eugene O''Neill', 'Thornton Wilder',
        'David Mamet', 'Neil Simon', 'Tony Kushner', 'August Wilson', 'Sam Shepard', 'Edward Albee',
        'Bob Fosse', 'Jerome Robbins', 'Agnes de Mille', 'Martha Graham', 'George Balanchine', 'Mikhail Baryshnikov',
        'Franco Dragone', 'Guy Laliberté', 'Robert Lepage', 'Julie Taymor', 'Mary Zimmerman', 'Des McAnuff'
    ];
BEGIN
    -- Get all location IDs
    SELECT ARRAY(SELECT id FROM locations ORDER BY random()) INTO location_ids;
    
    -- Get all theatre type IDs
    SELECT ARRAY(SELECT id FROM theatre_types ORDER BY random()) INTO theatre_type_ids;
    
    -- Get all show type IDs
    SELECT ARRAY(SELECT id FROM show_types ORDER BY random()) INTO show_type_ids;
    
    -- Insert 120 theaters
    FOR i IN 1..120 LOOP
        current_theatre_id := gen_random_uuid();
        
        INSERT INTO theatres (
            id, name, description, capacity, address, phone, email, website, image_url, 
            is_featured, is_active, location_id, theatre_type_id, created_at, updated_at
        ) VALUES (
            current_theatre_id,
            theatre_names[((i-1) % array_length(theatre_names, 1)) + 1] || CASE WHEN i > array_length(theatre_names, 1) THEN ' ' || (i / array_length(theatre_names, 1) + 1)::text ELSE '' END,
            'Professional theater venue offering world-class productions and entertainment',
            (random() * 2000 + 200)::integer,
            (random() * 999 + 100)::text || ' Theater Street, City',
            '(' || (random() * 900 + 100)::integer || ') ' || (random() * 900 + 100)::integer || '-' || (random() * 9000 + 1000)::integer,
            'info@theater' || i || '.com',
            'https://theater' || i || '.com',
            'https://example.com/theater' || i || '.jpg',
            (random() < 0.15), -- 15% chance of being featured
            true,
            location_ids[((i-1) % array_length(location_ids, 1)) + 1],
            theatre_type_ids[((i-1) % array_length(theatre_type_ids, 1)) + 1],
            NOW(),
            NOW()
        );
        
        -- Insert 2-4 shows per theater
        FOR j IN 1..(2 + (random() * 3)::integer) LOOP
            DECLARE
                show_start_date DATE;
                show_end_date DATE;
            BEGIN
                -- Generate start date within past 6 months to next 6 months
                show_start_date := CURRENT_DATE + (random() * 365 - 180)::integer;
                -- Generate end date 30-365 days after start date
                show_end_date := show_start_date + (random() * 335 + 30)::integer;
                
                INSERT INTO shows (
                    id, title, description, director, "cast", duration, start_date, end_date, 
                    price, image_url, trailer_url, is_featured, is_active, theatre_id, show_type_id, 
                    created_at, updated_at
                ) VALUES (
                    gen_random_uuid(),
                    show_titles[((i*j-1) % array_length(show_titles, 1)) + 1],
                    'A captivating theatrical experience that will leave audiences mesmerized with outstanding performances and stunning production values.',
                    directors[((i*j-1) % array_length(directors, 1)) + 1],
                    'Talented ensemble cast featuring award-winning performers',
                    (90 + random() * 120)::integer, -- Duration between 90-210 minutes
                    show_start_date,
                    show_end_date,
                    (25 + random() * 175)::numeric(10,2), -- Price between $25-$200
                    'https://example.com/show' || (i*100+j) || '.jpg',
                    CASE WHEN random() < 0.7 THEN 'https://youtube.com/watch?v=show' || (i*100+j) ELSE NULL END,
                    (random() < 0.1), -- 10% chance of being featured
                    true,
                    current_theatre_id,
                    show_type_ids[((i*j-1) % array_length(show_type_ids, 1)) + 1],
                    NOW(),
                    NOW()
                );
            END;
        END LOOP;
    END LOOP;
END $$;