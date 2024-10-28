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

INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate) VALUES
(1, '2024-10-20 04:00:00', 0.7, 98.1, 72),
(1, '2024-10-20 08:00:00', 0.6, 98.5, 73),
(1, '2024-10-20 12:00:00', 0.8, 97.8, 74),
(1, '2024-10-20 16:00:00', 0.9, 98.2, 72),
(1, '2024-10-20 20:00:00', 0.6, 98.9, 75),
(1, '2024-10-21 00:00:00', 0.7, 97.3, 70),
(1, '2024-10-21 04:00:00', 0.7, 99.1, 72),
(1, '2024-10-21 08:00:00', 0.8, 98.6, 73),
(1, '2024-10-21 12:00:00', 0.8, 98.8, 74),
(1, '2024-10-21 16:00:00', 0.6, 98.3, 75);

INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate) VALUES
(2, '2024-10-21 04:00:00', 1.0, 95.6, 78),
(2, '2024-10-21 08:00:00', 1.1, 96.1, 79),
(2, '2024-10-21 12:00:00', 1.2, 95.5, 80),
(2, '2024-10-21 16:00:00', 1.0, 96.8, 78),
(2, '2024-10-21 20:00:00', 1.1, 96.3, 81),
(2, '2024-10-22 00:00:00', 1.1, 95.9, 79),
(2, '2024-10-22 04:00:00', 1.3, 94.8, 82),
(2, '2024-10-22 08:00:00', 1.1, 96.4, 80),
(2, '2024-10-22 12:00:00', 1.0, 96.0, 77),
(2, '2024-10-22 16:00:00', 1.2, 95.3, 79);

INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate) VALUES
(3, '2024-10-22 04:00:00', 0.9, 97.2, 83),
(3, '2024-10-22 08:00:00', 0.8, 97.5, 82),
(3, '2024-10-22 12:00:00', 0.9, 97.1, 81),
(3, '2024-10-22 16:00:00', 0.8, 98.0, 84),
(3, '2024-10-22 20:00:00', 0.9, 97.4, 80),
(3, '2024-10-23 00:00:00', 0.9, 96.9, 85),
(3, '2024-10-23 04:00:00', 0.8, 97.3, 82),
(3, '2024-10-23 08:00:00', 0.9, 96.5, 81),
(3, '2024-10-23 12:00:00', 0.8, 98.1, 83),
(3, '2024-10-23 16:00:00', 0.7, 97.8, 80);

INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate) VALUES
(4, '2024-10-23 04:00:00', 1.3, 92.8, 88),
(4, '2024-10-23 08:00:00', 1.4, 93.0, 89),
(4, '2024-10-23 12:00:00', 1.5, 92.5, 90),
(4, '2024-10-23 16:00:00', 1.4, 92.7, 87),
(4, '2024-10-23 20:00:00', 1.5, 91.9, 91),
(4, '2024-10-24 00:00:00', 1.4, 93.2, 88),
(4, '2024-10-24 04:00:00', 1.3, 92.1, 87),
(4, '2024-10-24 08:00:00', 1.4, 91.8, 89),
(4, '2024-10-24 12:00:00', 1.3, 93.0, 86),
(4, '2024-10-24 16:00:00', 1.5, 92.9, 90);

INSERT INTO records (patient_id, test_date, bilirubin, oxygen, heart_rate) VALUES
(5, '2024-10-24 04:00:00', 0.6, 98.7, 69),
(5, '2024-10-24 08:00:00', 0.5, 99.0, 68),
(5, '2024-10-24 12:00:00', 0.6, 98.5, 67),
(5, '2024-10-24 16:00:00', 0.5, 98.9, 66),
(5, '2024-10-24 20:00:00', 0.6, 99.1, 69),
(5, '2024-10-25 00:00:00', 0.5, 99.2, 66),
(5, '2024-10-25 04:00:00', 0.6, 98.8, 68),
(5, '2024-10-25 08:00:00', 0.5, 98.9, 67),
(5, '2024-10-25 12:00:00', 0.6, 98.6, 69),
(5, '2024-10-25 16:00:00', 0.5, 99.0, 65);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE records;
-- +goose StatementEnd
