DROP TABLE IF EXISTS userinformation;
CREATE TABLE userinformation (
    id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    address VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT now() NOT NULL
);
