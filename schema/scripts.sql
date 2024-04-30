CREATE TABLE product (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE locale (
  id VARCHAR(36) PRIMARY KEY,
  locale VARCHAR(255)
);

CREATE TABLE language (
  id VARCHAR(36) PRIMARY KEY,
  language VARCHAR(255)
);

CREATE TABLE locale_language (
  id VARCHAR(36) PRIMARY KEY,
  locale_id VARCHAR(36),
  language_id VARCHAR(36),
  FOREIGN KEY (locale_id) REFERENCES locale(id),
  FOREIGN KEY (language_id) REFERENCES language(id)
);

CREATE TABLE consent_template (
  id VARCHAR(36) PRIMARY KEY,
  locale_language_id VARCHAR(36),
  template_pdf BLOB,
  FOREIGN KEY (locale_language_id) REFERENCES locale_language(id)
);

CREATE TABLE product_template (
  id VARCHAR(36) PRIMARY KEY, 
  product_id VARCHAR(36) NOT NULL,
  consent_template_id VARCHAR(36) NOT NULL,
  active BOOLEAN NOT NULL,
  FOREIGN KEY (consent_template_id) REFERENCES consent_template(id),
  FOREIGN KEY (product_id) REFERENCES product(id),
  INDEX idx_active_product (active, product_id)
);

CREATE TABLE client_consent (
  id VARCHAR(36) PRIMARY KEY,
  product_template_id VARCHAR(36),
  individual_id VARCHAR(255),
  date DATE,
  optout_reason VARCHAR(255),
  FOREIGN KEY (product_template_id) REFERENCES product_template(id)
);