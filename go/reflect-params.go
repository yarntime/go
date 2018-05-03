func initAppConfig(customConfig v1.CustomConfig, appConfig v1.ApplicationConfig) {
	globalFiled := reflect.TypeOf(customConfig.Global)
	globalValue := reflect.ValueOf(customConfig.Global)
	params := []string{}
	for i := 0; i < globalFiled.NumField(); i++ {
		f := globalFiled.Field(i)
		_, skip := f.Tag.Lookup("skip")
		if skip {
			continue
		}
		params = append(params, "--" + f.Tag.Get("json") + "=" + globalValue.Field(i).Interface().(string))
	}
	for i := 0; i < len(appConfig.App); i++ {
		appConfig.App[i].Params = append(appConfig.App[i].Params, params...)
	}
}