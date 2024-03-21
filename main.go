package main

import (
	"os"
	"fmt"
	"strings"
	host_client "github.com/matehaxor03/holistic_host_client/host_client"
	common "github.com/matehaxor03/holistic_common/common"
	host_installer "github.com/matehaxor03/holistic_host_init/host_installer"
	db_installer "github.com/matehaxor03/holistic_db_init/db_installer"
)

func main(){
	var errors []error
	host_client, host_client_errors := host_client.NewHostClient()
	if host_client_errors != nil {
		fmt.Println(fmt.Errorf("%s", host_client_errors))
		os.Exit(1)
	}

	number_of_users_value, number_of_users_errors := host_client.GetEnviornmentVariableValue(common.ENV_HOLISTIC_HOST_NUMBER_OF_USERS())
	if number_of_users_errors != nil {
		errors = append(errors, number_of_users_errors...)
	}

	users_offset_value, users_offset_value_errors := host_client.GetEnviornmentVariableValue(common.ENV_HOLISTIC_HOST_USERS_USERID_OFFSET())
	if users_offset_value_errors != nil {
		errors = append(errors, users_offset_value_errors...)
	}

	database_host_name, database_host_name_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_HOSTNAME())
	if database_host_name_errors != nil {
		errors = append(errors, database_host_name_errors...)
	}

	database_port_number, database_port_number_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_PORT_NUMBER())
	if database_port_number_errors != nil {
		errors = append(errors, database_port_number_errors...)
	}

	database_name, database_name_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_NAME())
	if database_name_errors != nil {
		errors = append(errors, database_name_errors...)
	}

	database_root_username, database_root_username_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_ROOT_USERNAME())
	if database_root_username_errors != nil {
		errors = append(errors, database_root_username_errors...)
	}

	database_root_password, database_root_password_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_ROOT_PASSWORD())
	if database_root_password_errors != nil {
		errors = append(errors, database_root_password_errors...)
	}


	database_writer_usernames, database_writer_usernames_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_WRITER_USERNAMES())
	if database_writer_usernames_errors != nil {
		errors = append(errors, database_writer_usernames_errors...)
	}

	database_reader_usernames, database_reader_usernames_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_READER_USERNAMES())
	if database_reader_usernames_errors != nil {
		errors = append(errors, database_reader_usernames_errors...)
	}

	database_migration_usernames, database_migration_usernames_errors := host_client.GetEnviornmentVariable(common.ENV_HOLISTIC_DATABASE_MIGRATION_USERNAMES())
	if database_migration_usernames_errors != nil {
		errors = append(errors, database_migration_usernames_errors...)
	}

	if len(errors) > 0 {
		fmt.Println(fmt.Errorf("%s", errors))
		os.Exit(1)
	}

	var database_writer_usernames_array []string
	var database_reader_usernames_array []string
	var database_migration_usernames_array []string

	if strings.Contains(*database_writer_usernames, ",") {
		database_writer_usernames_array = strings.Split(*database_writer_usernames, ",")
	} else {
		database_writer_usernames_array = append(database_writer_usernames_array, *database_writer_usernames)
	}

	if strings.Contains(*database_reader_usernames, ",") {
		database_reader_usernames_array = strings.Split(*database_reader_usernames, ",")
	} else {
		database_reader_usernames_array = append(database_reader_usernames_array, *database_reader_usernames)
	}

	if strings.Contains(*database_migration_usernames, ",") {
		database_migration_usernames_array = strings.Split(*database_migration_usernames, ",")
	} else {
		database_migration_usernames_array = append(database_migration_usernames_array, *database_migration_usernames)
	}

	number_of_users, number_of_users_uint64_errors := number_of_users_value.GetUInt64Value()
	if number_of_users_uint64_errors != nil {
		errors = append(errors, number_of_users_uint64_errors...)
	}

	userid_offset, userid_offset_uint64_errors := users_offset_value.GetUInt64Value()
	if userid_offset_uint64_errors != nil {
		errors = append(errors, userid_offset_uint64_errors...)
	}

	if len(errors) > 0 {
		fmt.Println(fmt.Errorf("%s", errors))
		os.Exit(1)
	}

	host_installer, host_installer_errors := host_installer.NewHostInstaller(number_of_users, userid_offset)
	if host_installer_errors != nil {
		errors = append(errors, host_installer_errors...)
	}

	database_installer,  database_installer_errors := db_installer.NewDatabaseInstaller(*database_host_name, *database_port_number, *database_name, *database_root_username, *database_root_password, database_writer_usernames_array, database_reader_usernames_array, database_migration_usernames_array)
	if database_installer_errors != nil {
		errors = append(errors, database_installer_errors...)
	}

	if len(errors) > 0 {
		fmt.Println(fmt.Errorf("%s", errors))
		os.Exit(1)
	}

	host_installer_install_errors := host_installer.Install()
	if host_installer_install_errors != nil {
		fmt.Println(fmt.Errorf("%s", host_installer_install_errors))
		os.Exit(1)
	}

	database_installer_install_errors := database_installer.Install()
	if database_installer_install_errors != nil {
		fmt.Println(fmt.Errorf("%s", database_installer_install_errors))
		os.Exit(1)
	}

	fmt.Println("install successfull")
	os.Exit(0)
}