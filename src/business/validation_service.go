package business

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ValidationService provides additional validation utilities
type ValidationService struct{}

// NewValidationService creates a new validation service
func NewValidationService() *ValidationService {
	return &ValidationService{}
}

// ValidateUUID validates if a string is a valid UUID
func (vs *ValidationService) ValidateUUID(uuidStr string) error {
	if uuidStr == "" {
		return errors.New("UUID cannot be empty")
	}

	_, err := uuid.Parse(uuidStr)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %s", uuidStr)
	}

	return nil
}

// ValidateEmail validates email format
func (vs *ValidationService) ValidateEmail(email string) error {
	if email == "" {
		return nil // Email is optional in most cases
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format: %s", email)
	}

	return nil
}

// ValidateURL validates URL format
func (vs *ValidationService) ValidateURL(url string) error {
	if url == "" {
		return nil // URL is optional in most cases
	}

	urlRegex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	if !urlRegex.MatchString(url) {
		return fmt.Errorf("invalid URL format: %s", url)
	}

	return nil
}

// ValidatePhoneNumber validates phone number format
func (vs *ValidationService) ValidatePhoneNumber(phone string) error {
	if phone == "" {
		return nil // Phone is optional in most cases
	}

	// Remove common separators and spaces
	cleanPhone := strings.ReplaceAll(phone, " ", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, "-", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, "(", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, ")", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, "+", "")

	// Check if it contains only digits and is reasonable length
	phoneRegex := regexp.MustCompile(`^\d{7,15}$`)
	if !phoneRegex.MatchString(cleanPhone) {
		return fmt.Errorf("invalid phone number format: %s", phone)
	}

	return nil
}

// ValidateCoordinates validates latitude and longitude
func (vs *ValidationService) ValidateCoordinates(latitude, longitude *float64) error {
	if latitude != nil {
		if *latitude < -90 || *latitude > 90 {
			return fmt.Errorf("latitude must be between -90 and 90, got: %f", *latitude)
		}
	}

	if longitude != nil {
		if *longitude < -180 || *longitude > 180 {
			return fmt.Errorf("longitude must be between -180 and 180, got: %f", *longitude)
		}
	}

	return nil
}

// ValidateDateRange validates that start date is before end date
func (vs *ValidationService) ValidateDateRange(startDate, endDate *time.Time) error {
	if startDate == nil || endDate == nil {
		return nil // One or both dates are optional
	}

	if endDate.Before(*startDate) {
		return fmt.Errorf("end date (%s) cannot be before start date (%s)",
			endDate.Format("2006-01-02"), startDate.Format("2006-01-02"))
	}

	return nil
}

// ValidateCapacity validates theatre capacity
func (vs *ValidationService) ValidateCapacity(capacity *int) error {
	if capacity == nil {
		return nil // Capacity is optional
	}

	if *capacity < 1 {
		return fmt.Errorf("capacity must be at least 1, got: %d", *capacity)
	}

	if *capacity > 100000 {
		return fmt.Errorf("capacity cannot exceed 100,000, got: %d", *capacity)
	}

	return nil
}

// ValidateDuration validates show duration in minutes
func (vs *ValidationService) ValidateDuration(duration *int) error {
	if duration == nil {
		return nil // Duration is optional
	}

	if *duration < 1 {
		return fmt.Errorf("duration must be at least 1 minute, got: %d", *duration)
	}

	if *duration > 600 { // 10 hours max
		return fmt.Errorf("duration cannot exceed 600 minutes (10 hours), got: %d", *duration)
	}

	return nil
}

// ValidatePrice validates price value
func (vs *ValidationService) ValidatePrice(price *float64) error {
	if price == nil {
		return nil // Price is optional
	}

	if *price < 0 {
		return fmt.Errorf("price cannot be negative, got: %f", *price)
	}

	if *price > 10000 { // $10,000 max reasonable price
		return fmt.Errorf("price seems unreasonably high, got: %f", *price)
	}

	return nil
}

// ValidateStringLength validates string length constraints
func (vs *ValidationService) ValidateStringLength(value, fieldName string, minLength, maxLength int) error {
	if len(value) < minLength {
		return fmt.Errorf("%s must be at least %d characters long, got %d", fieldName, minLength, len(value))
	}

	if len(value) > maxLength {
		return fmt.Errorf("%s cannot exceed %d characters, got %d", fieldName, maxLength, len(value))
	}

	return nil
}

// ValidatePostalCode validates postal code format (basic validation)
func (vs *ValidationService) ValidatePostalCode(postalCode string) error {
	if postalCode == "" {
		return nil // Postal code is optional
	}

	// Basic validation - alphanumeric with spaces and dashes
	postalRegex := regexp.MustCompile(`^[A-Za-z0-9\s\-]{3,10}$`)
	if !postalRegex.MatchString(postalCode) {
		return fmt.Errorf("invalid postal code format: %s", postalCode)
	}

	return nil
}
