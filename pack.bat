@REM this is done with paketo (same build process as fly io)
pack build golang-webapp --buildpack gcr.io/paketo-buildpacks/go --builder paketobuildpacks/builder:base
