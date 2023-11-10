-- insert_test_data.sql

INSERT INTO owners (owner_name, contact_info, date_of_birth)
VALUES
    ('Igor Youriévitch Bogdanoff', 'yesterday.in@norwegian.com', '1949-08-29'),
    ('Grichka Youriévitch Bogdanoff', 'griskhet@tvilling.com', '1949-08-29'),
    ('Ernst Snømann', 'ernst@snowmail.com', '1530-09-10');

INSERT INTO cars (make, model, year, vin, color, purchase_date, owner_id)
VALUES
    ('Toyota', 'Camry', 2019, '1HGCM82633A123456', 'Blue', '2020-03-15', 1),
    ('Ford', 'Focus', 2018, '2FADP4EJ6AJ123456', 'Red', '2020-07-22', 2),
    ('Honda', 'Civic', 2020, '19XFC2F54LE123456', 'Silver', '2021-01-10', 1);

INSERT INTO engines (car_id, type, displacement, horsepower, manufacturing_date)
VALUES
    (1, 'V6', '3.5L', 270, '2019-05-01'),
    (2, 'Inline-4', '2.0L', 160, '2018-12-15'),
    (3, 'Inline-4', '1.8L', 140, '2019-02-20');

INSERT INTO tires (car_id, brand, size, manufacturing_date)
VALUES
    (1, 'Bridgestone', '215/55R17', '2019-05-01'),
    (2, 'Michelin', '205/60R16', '2018-12-15'),
    (3, 'Goodyear', '195/65R15', '2019-02-20');

INSERT INTO repairs (car_id, repair_type, repair_date, cost)
VALUES
    (1, 'Oil Change', '2021-02-15', 50.00),
    (2, 'Brake Replacement', '2020-09-30', 120.00),
    (3, 'Tire Rotation', '2021-03-05', 25.00);

INSERT INTO insurance (car_id, policy_number, provider, start_date, end_date)
VALUES
    (1, 'ABC123', 'Allstate', '2021-04-01', '2022-04-01'),
    (2, 'XYZ789', 'Geico', '2020-12-15', '2021-12-15'),
    (3, 'PQR456', 'State Farm', '2021-02-10', '2022-02-10');
