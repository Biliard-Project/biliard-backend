-- +goose Up
-- +goose StatementBegin
CREATE TABLE patients(
	id serial primary key,
	name text,
	gender text,
	birth_date date
);

INSERT INTO patients (name, gender, birth_date)
VALUES
('John Doe', 'Male', '1990-01-15'),
('Jane Smith', 'Female', '1985-06-20'),
('Robert Johnson', 'Male', '1978-03-10'),
('Emily Davis', 'Female', '1992-11-05'),
('Michael Brown', 'Male', '1983-09-22'),
('Sarah Wilson', 'Female', '1991-12-30'),
('David Taylor', 'Male', '1987-08-14'),
('Laura Martin', 'Female', '1995-05-09'),
('James Anderson', 'Male', '1980-04-25'),
('Anna Thompson', 'Female', '1993-07-17');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE patients;
-- +goose StatementEnd
