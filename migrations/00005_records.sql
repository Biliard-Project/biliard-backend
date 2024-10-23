-- +goose Up
-- +goose StatementBegin
CREATE TABLE records(
	id serial primary key,
	patient_id int not null references patients (id) on delete cascade,
	test_date timestamp,
	bilirubin float,
	oxygen float,
	heart_rate float
);

INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate)
VALUES
(1, '2024-10-20 08:00:00', 0.6, 98.8, 73),
(2, '2024-10-21 12:00:00', 1.0, 96.5, 77),
(3, '2024-10-22 16:00:00', 0.8, 97.8, 82),
(4, '2024-10-23 08:00:00', 1.4, 92.3, 89),
(5, '2024-10-24 12:00:00', 0.7, 98.9, 70),
(1, '2024-10-25 16:00:00', 0.9, 98.2, 74),
(3, '2024-10-26 08:00:00', 0.9, 96.7, 81),
(2, '2024-10-27 12:00:00', 1.1, 95.8, 79),
(4, '2024-10-28 16:00:00', 1.3, 93.0, 87),
(5, '2024-10-29 08:00:00', 0.5, 99.2, 66);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE records;
-- +goose StatementEnd
