package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/keikoproj/alert-manager/pkg/log"
	"strings"
)

//ExportParamsChecksum function calculates checksum if exportParams is not empty
func ExportParamsChecksum(ctx context.Context, exportedParams []string) (bool, string) {
	log := log.Logger(ctx, "utils", "util", "ExportParamsChecksum")
	if len(exportedParams) == 0 {
		return false, ""
	}
	log.V(4).Info("exportedParams are not empty")
	return true, calculateChecksum(ctx, strings.Join(exportedParams, ""))
}

//calculateChecksum function calculates checksum for the given string
func calculateChecksum(ctx context.Context, input string) string {
	log := log.Logger(ctx, "utils", "util", "calculateChecksum")
	log.V(4).Info("calculating checksum", "input", input)
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

//ContainsString  Helper functions to check from a slice of strings.
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

//RemoveString Helper function to check remove string
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}