-- create_tables.sql

CREATE TABLE IF NOT EXISTS owners (
                                      id SERIAL PRIMARY KEY,
                                      owner_name TEXT,
                                      contact_info TEXT,
                                      date_of_birth DATE
);

CREATE TABLE IF NOT EXISTS cars (
                                    id SERIAL PRIMARY KEY,
                                    make TEXT,
                                    model TEXT,
                                    year INT,
                                    vin TEXT UNIQUE,
                                    color TEXT,
                                    purchase_date DATE,
                                    owner_id INT REFERENCES owners(id)
);

CREATE TABLE IF NOT EXISTS engines (
                                       id SERIAL PRIMARY KEY,
                                       car_id INT REFERENCES cars(id),
                                       type TEXT,
                                       displacement TEXT,
                                       horsepower INT,
                                       manufacturing_date DATE
);

CREATE TABLE IF NOT EXISTS tires (
                                     id SERIAL PRIMARY KEY,
                                     car_id INT REFERENCES cars(id),
                                     brand TEXT,
                                     size TEXT,
                                     manufacturing_date DATE
);

CREATE TABLE IF NOT EXISTS repairs (
                                       id SERIAL PRIMARY KEY,
                                       car_id INT REFERENCES cars(id),
                                       repair_type TEXT,
                                       repair_date DATE,
                                       cost DECIMAL(10, 2)
);

CREATE TABLE IF NOT EXISTS insurances (
                                         id SERIAL PRIMARY KEY,
                                         car_id INT REFERENCES cars(id),
                                         policy_number TEXT,
                                         provider TEXT,
                                         start_date DATE,
                                         end_date DATE
);
