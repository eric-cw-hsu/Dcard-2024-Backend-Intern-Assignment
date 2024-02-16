package repositories

import (
	"context"
	"database/sql"
	"dcard-2024-backend-intern-assignment/models"
	"time"
)

type AdRepository struct {
	db *sql.DB
}

func NewAdRepository(db *sql.DB) *AdRepository {
	return &AdRepository{db}
}

func (r *AdRepository) CreateAd(data models.Ad) (models.Ad, error) {
	var ad models.Ad
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return ad, err
	}

	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, `
		INSERT INTO ads (
			title, start_at, end_at, age_start, age_end, gender
		) VALUES (
			?, ?, ?, ?, ?, ?
		);
	`, data.Title,
		data.StartAt, data.EndAt, data.Conditions.AgeStart,
		data.Conditions.AgeEnd, data.Conditions.Gender,
	)

	if err != nil {
		return ad, err
	}

	adId, _ := result.LastInsertId()

	// create ad condition platform
	err = r.createAdConditionPlatform(ctx, tx, adId, data.Conditions.Platform)
	if err != nil {
		return ad, err
	}

	// create ad condition country
	err = r.createAdConditionCountry(ctx, tx, adId, data.Conditions.Country)
	if err != nil {
		return ad, err
	}

	if err = tx.Commit(); err != nil {
		return ad, err
	}

	ad, err = r.GetAdById(adId)
	if err != nil {
		return ad, err
	}

	return ad, nil
}

func (r *AdRepository) GetAdById(id int64) (models.Ad, error) {
	var ad models.Ad
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.db.QueryRowContext(ctx, `
		SELECT id, title, start_at, end_at, age_start, age_end, gender
		FROM ads
		WHERE id = ?;
	`, id).Scan(&ad.Id, &ad.Title, &ad.StartAt, &ad.EndAt, &ad.Conditions.AgeStart, &ad.Conditions.AgeEnd, &ad.Conditions.Gender)

	if err != nil {
		return ad, err
	}

	// get ad condition platform
	rows, err := r.db.QueryContext(ctx, `
		SELECT platform
		FROM ads_to_platforms
		WHERE ad_id = ?;
	`, id)

	if err != nil {
		return ad, err
	}

	for rows.Next() {
		var platform string
		err = rows.Scan(&platform)
		if err != nil {
			return ad, err
		}

		ad.Conditions.Platform = append(ad.Conditions.Platform, platform)
	}

	// get ad condition country
	rows, err = r.db.QueryContext(ctx, `
		SELECT country_code
		FROM ads_to_country_codes
		WHERE ad_id = ?;
	`, id)

	if err != nil {
		return ad, err
	}

	for rows.Next() {
		var country string
		err = rows.Scan(&country)
		if err != nil {
			return ad, err
		}

		ad.Conditions.Country = append(ad.Conditions.Country, country)
	}

	return ad, nil
}

func (r *AdRepository) createAdConditionPlatform(ctx context.Context, tx *sql.Tx, adId int64, platforms []string) error {
	// find the platform id
	for _, platform := range platforms {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO ads_to_platforms (ad_id, platform)
			VALUES (?, ?);
		`, adId, platform)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *AdRepository) createAdConditionCountry(ctx context.Context, tx *sql.Tx, adId int64, countries []string) error {
	for _, country_code := range countries {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO ads_to_country_codes (ad_id, country_code)
			VALUES (?, ?);
		`, adId, country_code)

		if err != nil {
			return err
		}
	}

	return nil
}
