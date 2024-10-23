-- +goose Up
-- +goose StatementBegin
CREATE TABLE patients(
	id serial primary key,
	name text,
	gender text,
	birth_date date,
	keterangan text
);

INSERT INTO patients (name, gender, birth_date, keterangan)
VALUES
('John Doe', 'Male', '1985-02-15', 'No known allergies'),
('Jane Smith', 'Female', '1990-06-25', 'Diabetic'),
('Alice Brown', 'Female', '1975-11-05', 'High blood pressure'),
('Bob Johnson', 'Male', '1965-04-20', 'Smoker, history of heart disease'),
('Eve Wilson', 'Female', '2001-08-13', 'Healthy');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE patients;
-- +goose StatementEnd
