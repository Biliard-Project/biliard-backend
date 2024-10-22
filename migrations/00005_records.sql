-- +goose Up
-- +goose StatementBegin
CREATE TABLE records(
	id serial primary key,
	patient_id int not null references patients (id) on delete cascade,
	keterangan text,
	test_date date
);

INSERT INTO records (patient_id, keterangan, test_date)
VALUES
(1, 'normal', '2024-01-10'),
(2, 'rendah', '2024-01-15'),
(3, 'tinggi', '2024-01-20'),
(4, 'normal', '2024-01-22'),
(5, 'tinggi', '2024-01-25'),
(6, 'rendah', '2024-01-28'),
(7, 'normal', '2024-02-01'),
(8, 'tinggi', '2024-02-05'),
(9, 'rendah', '2024-02-10'),
(10, 'normal', '2024-02-15');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE records;
-- +goose StatementEnd
