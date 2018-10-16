package darksky

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const response = `{"latitude":53.5458874,"longitude":-113.5034304,"timezone":"America/Edmonton","currently":{"time":1539463587,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":32.94,"apparentTemperature":24.79,"dewPoint":22.31,"humidity":0.65,"pressure":1028.37,"windSpeed":10.21,"windGust":17.91,"windBearing":337,"cloudCover":0.92,"uvIndex":1,"visibility":10,"ozone":307.15},"hourly":{"summary":"Mostly cloudy starting tonight.","icon":"partly-cloudy-day","data":[{"time":1539460800,"summary":"Overcast","icon":"cloudy","precipIntensity":0,"precipProbability":0,"temperature":32.12,"apparentTemperature":23.71,"dewPoint":21.64,"humidity":0.65,"pressure":1028.41,"windSpeed":10.33,"windGust":18.6,"windBearing":338,"cloudCover":1,"uvIndex":2,"visibility":10,"ozone":306.34},{"time":1539464400,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":33.18,"apparentTemperature":25.1,"dewPoint":22.51,"humidity":0.64,"pressure":1028.36,"windSpeed":10.18,"windGust":17.7,"windBearing":337,"cloudCover":0.89,"uvIndex":1,"visibility":10,"ozone":307.38},{"time":1539468000,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":34.55,"apparentTemperature":27.08,"dewPoint":23.64,"humidity":0.64,"pressure":1028.52,"windSpeed":9.61,"windGust":16.1,"windBearing":335,"cloudCover":0.67,"uvIndex":1,"visibility":10,"ozone":308.2},{"time":1539471600,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":35.88,"apparentTemperature":29.28,"dewPoint":24.83,"humidity":0.64,"pressure":1028.83,"windSpeed":8.53,"windGust":13.94,"windBearing":333,"cloudCover":0.46,"uvIndex":0,"visibility":10,"ozone":308.77},{"time":1539475200,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0003,"precipProbability":0.06,"precipAccumulation":0,"precipType":"snow","temperature":36.08,"apparentTemperature":30.32,"dewPoint":25.34,"humidity":0.65,"pressure":1029.14,"windSpeed":7.16,"windGust":11.47,"windBearing":334,"cloudCover":0.34,"uvIndex":0,"visibility":10,"ozone":309.3},{"time":1539478800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":35.12,"apparentTemperature":30.34,"dewPoint":24.91,"humidity":0.66,"pressure":1029.52,"windSpeed":5.48,"windGust":8.38,"windBearing":325,"cloudCover":0.28,"uvIndex":0,"visibility":10,"ozone":309.63},{"time":1539482400,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":33.12,"apparentTemperature":29.58,"dewPoint":24.13,"humidity":0.69,"pressure":1029.95,"windSpeed":3.81,"windGust":5.03,"windBearing":266,"cloudCover":0.19,"uvIndex":0,"visibility":10,"ozone":309.82},{"time":1539486000,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":31.73,"apparentTemperature":31.73,"dewPoint":23.55,"humidity":0.71,"pressure":1030.22,"windSpeed":2.79,"windGust":2.84,"windBearing":280,"cloudCover":0.12,"uvIndex":0,"visibility":10,"ozone":310.01},{"time":1539489600,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":31.15,"apparentTemperature":31.15,"dewPoint":23.39,"humidity":0.73,"pressure":1030.28,"windSpeed":2.84,"windGust":2.87,"windBearing":3,"cloudCover":0.11,"uvIndex":0,"visibility":10,"ozone":310.44},{"time":1539493200,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":30.84,"apparentTemperature":27.25,"dewPoint":23.37,"humidity":0.73,"pressure":1030.22,"windSpeed":3.57,"windGust":4.05,"windBearing":157,"cloudCover":0.1,"uvIndex":0,"visibility":10,"ozone":311.01},{"time":1539496800,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":30.75,"apparentTemperature":26.36,"dewPoint":23.55,"humidity":0.74,"pressure":1030.16,"windSpeed":4.25,"windGust":5.18,"windBearing":230,"cloudCover":0.12,"uvIndex":0,"visibility":10,"ozone":310.56},{"time":1539500400,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":30.89,"apparentTemperature":26.12,"dewPoint":24.03,"humidity":0.75,"pressure":1030.05,"windSpeed":4.64,"windGust":5.74,"windBearing":244,"cloudCover":0.17,"uvIndex":0,"visibility":10,"ozone":308.62},{"time":1539504000,"summary":"Clear","icon":"clear-night","precipIntensity":0,"precipProbability":0,"temperature":31.43,"apparentTemperature":26.42,"dewPoint":24.64,"humidity":0.76,"pressure":1029.93,"windSpeed":4.99,"windGust":6.2,"windBearing":233,"cloudCover":0.24,"uvIndex":0,"visibility":10,"ozone":305.79},{"time":1539507600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0002,"precipProbability":0.02,"precipAccumulation":0,"precipType":"snow","temperature":31.63,"apparentTemperature":26.37,"dewPoint":25.03,"humidity":0.76,"pressure":1029.77,"windSpeed":5.3,"windGust":6.67,"windBearing":222,"cloudCover":0.29,"uvIndex":0,"visibility":10,"ozone":303.29},{"time":1539511200,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":31.36,"apparentTemperature":25.89,"dewPoint":25.03,"humidity":0.77,"pressure":1029.5,"windSpeed":5.48,"windGust":7.24,"windBearing":223,"cloudCover":0.4,"uvIndex":0,"visibility":10,"ozone":301.44},{"time":1539514800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":29.93,"apparentTemperature":24.12,"dewPoint":24.76,"humidity":0.81,"pressure":1029.18,"windSpeed":5.57,"windGust":7.8,"windBearing":224,"cloudCover":0.53,"uvIndex":0,"visibility":9.29,"ozone":299.82},{"time":1539518400,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":29.43,"apparentTemperature":23.45,"dewPoint":24.51,"humidity":0.82,"pressure":1028.86,"windSpeed":5.67,"windGust":8.18,"windBearing":224,"cloudCover":0.64,"uvIndex":0,"visibility":6.97,"ozone":298.29},{"time":1539522000,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":28.46,"apparentTemperature":22.26,"dewPoint":24.06,"humidity":0.83,"pressure":1028.6,"windSpeed":5.71,"windGust":7.96,"windBearing":222,"cloudCover":0.7,"uvIndex":0,"visibility":6.84,"ozone":296.52},{"time":1539525600,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":27.91,"apparentTemperature":21.58,"dewPoint":23.65,"humidity":0.84,"pressure":1028.36,"windSpeed":5.75,"windGust":7.52,"windBearing":220,"cloudCover":0.78,"uvIndex":0,"visibility":6.73,"ozone":294.83},{"time":1539529200,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":29.14,"apparentTemperature":22.81,"dewPoint":23.84,"humidity":0.8,"pressure":1027.84,"windSpeed":6.02,"windGust":7.98,"windBearing":219,"cloudCover":0.84,"uvIndex":0,"visibility":7.63,"ozone":293.37},{"time":1539532800,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":32.66,"apparentTemperature":26.42,"dewPoint":25.17,"humidity":0.74,"pressure":1026.97,"windSpeed":6.82,"windGust":10.28,"windBearing":224,"cloudCover":0.74,"uvIndex":0,"visibility":9.86,"ozone":292.04},{"time":1539536400,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":38.1,"apparentTemperature":32.36,"dewPoint":27.15,"humidity":0.64,"pressure":1025.8,"windSpeed":7.85,"windGust":13.32,"windBearing":232,"cloudCover":0.59,"uvIndex":1,"visibility":10,"ozone":290.93},{"time":1539540000,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":42.74,"apparentTemperature":37.64,"dewPoint":29.03,"humidity":0.58,"pressure":1024.66,"windSpeed":8.6,"windGust":15.34,"windBearing":231,"cloudCover":0.5,"uvIndex":2,"visibility":10,"ozone":290.16},{"time":1539543600,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":45.71,"apparentTemperature":41.18,"dewPoint":30.42,"humidity":0.55,"pressure":1023.59,"windSpeed":8.84,"windGust":15.55,"windBearing":204,"cloudCover":0.55,"uvIndex":2,"visibility":10,"ozone":289.84},{"time":1539547200,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":47.49,"apparentTemperature":43.36,"dewPoint":31.54,"humidity":0.54,"pressure":1022.52,"windSpeed":8.84,"windGust":14.82,"windBearing":243,"cloudCover":0.66,"uvIndex":2,"visibility":10,"ozone":289.72},{"time":1539550800,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0003,"precipProbability":0.04,"precipType":"rain","temperature":48.42,"apparentTemperature":44.57,"dewPoint":32.49,"humidity":0.54,"pressure":1021.57,"windSpeed":8.68,"windGust":13.95,"windBearing":243,"cloudCover":0.73,"uvIndex":1,"visibility":10,"ozone":289.26},{"time":1539554400,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0.0002,"precipProbability":0.04,"precipType":"rain","temperature":47.99,"apparentTemperature":44.24,"dewPoint":33.38,"humidity":0.57,"pressure":1020.69,"windSpeed":8.18,"windGust":12.77,"windBearing":255,"cloudCover":0.68,"uvIndex":1,"visibility":10,"ozone":287.93},{"time":1539558000,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":47.09,"apparentTemperature":43.43,"dewPoint":34.14,"humidity":0.6,"pressure":1019.94,"windSpeed":7.54,"windGust":11.44,"windBearing":221,"cloudCover":0.59,"uvIndex":0,"visibility":10,"ozone":286.23},{"time":1539561600,"summary":"Partly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":46.16,"apparentTemperature":42.42,"dewPoint":34.64,"humidity":0.64,"pressure":1019.42,"windSpeed":7.3,"windGust":11.32,"windBearing":245,"cloudCover":0.53,"uvIndex":0,"visibility":10,"ozone":285.06},{"time":1539565200,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":45.31,"apparentTemperature":41.15,"dewPoint":34.86,"humidity":0.67,"pressure":1019.27,"windSpeed":7.79,"windGust":13.06,"windBearing":235,"cloudCover":0.5,"uvIndex":0,"visibility":10,"ozone":284.73},{"time":1539568800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":44.47,"apparentTemperature":39.72,"dewPoint":34.82,"humidity":0.69,"pressure":1019.37,"windSpeed":8.69,"windGust":15.93,"windBearing":276,"cloudCover":0.5,"uvIndex":0,"visibility":10,"ozone":284.85},{"time":1539572400,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":43.87,"apparentTemperature":38.54,"dewPoint":34.43,"humidity":0.69,"pressure":1019.44,"windSpeed":9.74,"windGust":19.25,"windBearing":250,"cloudCover":0.51,"uvIndex":0,"visibility":10,"ozone":284.96},{"time":1539576000,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":43.49,"apparentTemperature":37.59,"dewPoint":33.45,"humidity":0.67,"pressure":1019.27,"windSpeed":10.99,"windGust":23.37,"windBearing":196,"cloudCover":0.55,"uvIndex":0,"visibility":10,"ozone":285.17},{"time":1539579600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":43.38,"apparentTemperature":36.93,"dewPoint":32.08,"humidity":0.64,"pressure":1019.07,"windSpeed":12.49,"windGust":28.01,"windBearing":329,"cloudCover":0.59,"uvIndex":0,"visibility":10,"ozone":285.4},{"time":1539583200,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0.0002,"precipProbability":0.02,"precipType":"rain","temperature":43.5,"apparentTemperature":36.73,"dewPoint":30.87,"humidity":0.61,"pressure":1018.99,"windSpeed":13.57,"windGust":31.6,"windBearing":275,"cloudCover":0.59,"uvIndex":0,"visibility":10,"ozone":285.17},{"time":1539586800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":43.08,"apparentTemperature":36.07,"dewPoint":30.04,"humidity":0.6,"pressure":1019.17,"windSpeed":14,"windGust":33.63,"windBearing":258,"cloudCover":0.49,"uvIndex":0,"visibility":10,"ozone":284.17},{"time":1539590400,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":42.6,"apparentTemperature":35.45,"dewPoint":29.41,"humidity":0.59,"pressure":1019.46,"windSpeed":14.02,"windGust":34.63,"windBearing":292,"cloudCover":0.34,"uvIndex":0,"visibility":10,"ozone":282.84},{"time":1539594000,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":42.13,"apparentTemperature":34.9,"dewPoint":29.01,"humidity":0.59,"pressure":1019.81,"windSpeed":13.87,"windGust":34.74,"windBearing":285,"cloudCover":0.26,"uvIndex":0,"visibility":10,"ozone":281.78},{"time":1539597600,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":41.58,"apparentTemperature":34.31,"dewPoint":28.96,"humidity":0.61,"pressure":1020.14,"windSpeed":13.55,"windGust":33.61,"windBearing":278,"cloudCover":0.28,"uvIndex":0,"visibility":10,"ozone":281.08},{"time":1539601200,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":41.16,"apparentTemperature":33.94,"dewPoint":29.18,"humidity":0.62,"pressure":1020.5,"windSpeed":13.05,"windGust":31.62,"windBearing":279,"cloudCover":0.37,"uvIndex":0,"visibility":10,"ozone":280.66},{"time":1539604800,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":40.11,"apparentTemperature":32.73,"dewPoint":29.29,"humidity":0.65,"pressure":1020.84,"windSpeed":12.68,"windGust":30,"windBearing":282,"cloudCover":0.47,"uvIndex":0,"visibility":10,"ozone":280.26},{"time":1539608400,"summary":"Partly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":38.94,"apparentTemperature":31.35,"dewPoint":29.17,"humidity":0.68,"pressure":1021.25,"windSpeed":12.43,"windGust":28.93,"windBearing":303,"cloudCover":0.56,"uvIndex":0,"visibility":10,"ozone":280.03},{"time":1539612000,"summary":"Mostly Cloudy","icon":"partly-cloudy-night","precipIntensity":0,"precipProbability":0,"temperature":38.29,"apparentTemperature":30.57,"dewPoint":28.91,"humidity":0.69,"pressure":1021.64,"windSpeed":12.31,"windGust":28.18,"windBearing":260,"cloudCover":0.66,"uvIndex":0,"visibility":10,"ozone":279.93},{"time":1539615600,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":39.32,"apparentTemperature":31.79,"dewPoint":29.12,"humidity":0.67,"pressure":1021.82,"windSpeed":12.53,"windGust":28.26,"windBearing":276,"cloudCover":0.73,"uvIndex":0,"visibility":10,"ozone":279.93},{"time":1539619200,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":42.24,"apparentTemperature":35.23,"dewPoint":30.05,"humidity":0.62,"pressure":1021.63,"windSpeed":13.29,"windGust":29.94,"windBearing":264,"cloudCover":0.74,"uvIndex":0,"visibility":10,"ozone":280.07},{"time":1539622800,"summary":"Mostly Cloudy","icon":"partly-cloudy-day","precipIntensity":0,"precipProbability":0,"temperature":46.46,"apparentTemperature":40.28,"dewPoint":31.46,"humidity":0.56,"pressure":1021.23,"windSpeed":14.34,"windGust":32.31,"windBearing":306,"cloudCover":0.72,"uvIndex":1,"visibility":10,"ozone":280.31},{"time":1539626400,"summary":"Breezy and Mostly Cloudy","icon":"wind","precipIntensity":0,"precipProbability":0,"temperature":49.63,"apparentTemperature":44.11,"dewPoint":32.82,"humidity":0.52,"pressure":1020.92,"windSpeed":15.24,"windGust":33.47,"windBearing":292,"cloudCover":0.69,"uvIndex":2,"visibility":10,"ozone":280.42},{"time":1539630000,"summary":"Breezy and Mostly Cloudy","icon":"wind","precipIntensity":0,"precipProbability":0,"temperature":52.27,"apparentTemperature":52.27,"dewPoint":34.04,"humidity":0.5,"pressure":1020.72,"windSpeed":15.9,"windGust":32.58,"windBearing":291,"cloudCover":0.63,"uvIndex":2,"visibility":10,"ozone":280.09},{"time":1539633600,"summary":"Breezy and Partly Cloudy","icon":"wind","precipIntensity":0,"precipProbability":0,"temperature":54.1,"apparentTemperature":54.1,"dewPoint":35.21,"humidity":0.49,"pressure":1020.56,"windSpeed":16.3,"windGust":30.63,"windBearing":298,"cloudCover":0.54,"uvIndex":2,"visibility":10,"ozone":279.66}]},"daily":{"summary":"No precipitation throughout the week, with high temperatures rising to 65°F on Wednesday.","icon":"clear-day","data":[{"time":1539410400,"summary":"Mostly cloudy until evening.","icon":"partly-cloudy-day","sunriseTime":1539439146,"sunsetTime":1539477806,"moonPhase":0.16,"precipIntensity":0.0002,"precipIntensityMax":0.0009,"precipIntensityMaxTime":1539410400,"precipProbability":0.24,"precipAccumulation":0.051,"precipType":"snow","temperatureHigh":36.08,"temperatureHighTime":1539475200,"temperatureLow":27.91,"temperatureLowTime":1539525600,"apparentTemperatureHigh":30.34,"apparentTemperatureHighTime":1539478800,"apparentTemperatureLow":21.58,"apparentTemperatureLowTime":1539525600,"dewPoint":22.72,"humidity":0.71,"pressure":1027.06,"windSpeed":7.17,"windGust":20.86,"windGustTime":1539414000,"windBearing":328,"cloudCover":0.74,"uvIndex":2,"uvIndexTime":1539457200,"visibility":10,"ozone":302.94,"temperatureMin":27.55,"temperatureMinTime":1539435600,"temperatureMax":36.08,"temperatureMaxTime":1539475200,"apparentTemperatureMin":19.16,"apparentTemperatureMinTime":1539435600,"apparentTemperatureMax":31.73,"apparentTemperatureMaxTime":1539486000},{"time":1539496800,"summary":"Partly cloudy throughout the day.","icon":"partly-cloudy-day","sunriseTime":1539525657,"sunsetTime":1539564067,"moonPhase":0.19,"precipIntensity":0.0001,"precipIntensityMax":0.0003,"precipIntensityMaxTime":1539550800,"precipProbability":0.07,"precipType":"rain","temperatureHigh":48.42,"temperatureHighTime":1539550800,"temperatureLow":38.29,"temperatureLowTime":1539612000,"apparentTemperatureHigh":44.57,"apparentTemperatureHighTime":1539550800,"apparentTemperatureLow":30.57,"apparentTemperatureLowTime":1539612000,"dewPoint":28.78,"humidity":0.69,"pressure":1024.74,"windSpeed":6.41,"windGust":28.01,"windGustTime":1539579600,"windBearing":236,"cloudCover":0.54,"uvIndex":2,"uvIndexTime":1539540000,"visibility":10,"ozone":293.28,"temperatureMin":27.91,"temperatureMinTime":1539525600,"temperatureMax":48.42,"temperatureMaxTime":1539550800,"apparentTemperatureMin":21.58,"apparentTemperatureMinTime":1539525600,"apparentTemperatureMax":44.57,"apparentTemperatureMaxTime":1539550800},{"time":1539583200,"summary":"Partly cloudy throughout the day and breezy in the afternoon.","icon":"wind","sunriseTime":1539612168,"sunsetTime":1539650329,"moonPhase":0.22,"precipIntensity":0.0001,"precipIntensityMax":0.0002,"precipIntensityMaxTime":1539583200,"precipProbability":0.15,"precipType":"rain","temperatureHigh":55.21,"temperatureHighTime":1539637200,"temperatureLow":37.17,"temperatureLowTime":1539698400,"apparentTemperatureHigh":55.21,"apparentTemperatureHighTime":1539637200,"apparentTemperatureLow":30.37,"apparentTemperatureLowTime":1539698400,"dewPoint":32.35,"humidity":0.58,"pressure":1021.35,"windSpeed":11.97,"windGust":34.74,"windGustTime":1539594000,"windBearing":288,"cloudCover":0.53,"uvIndex":2,"uvIndexTime":1539626400,"visibility":10,"ozone":280.42,"temperatureMin":38.29,"temperatureMinTime":1539612000,"temperatureMax":55.21,"temperatureMaxTime":1539637200,"apparentTemperatureMin":30.57,"apparentTemperatureMinTime":1539612000,"apparentTemperatureMax":55.21,"apparentTemperatureMaxTime":1539637200},{"time":1539669600,"summary":"Partly cloudy in the morning.","icon":"partly-cloudy-night","sunriseTime":1539698679,"sunsetTime":1539736592,"moonPhase":0.25,"precipIntensity":0.0001,"precipIntensityMax":0.0003,"precipIntensityMaxTime":1539684000,"precipProbability":0.05,"precipType":"rain","temperatureHigh":59.8,"temperatureHighTime":1539723600,"temperatureLow":40.59,"temperatureLowTime":1539784800,"apparentTemperatureHigh":59.8,"apparentTemperatureHighTime":1539723600,"apparentTemperatureLow":36.1,"apparentTemperatureLowTime":1539784800,"dewPoint":30.23,"humidity":0.52,"pressure":1025.05,"windSpeed":7.91,"windGust":18.66,"windGustTime":1539716400,"windBearing":269,"cloudCover":0.27,"uvIndex":2,"uvIndexTime":1539712800,"visibility":10,"ozone":284.42,"temperatureMin":37.17,"temperatureMinTime":1539698400,"temperatureMax":59.8,"temperatureMaxTime":1539723600,"apparentTemperatureMin":30.37,"apparentTemperatureMinTime":1539698400,"apparentTemperatureMax":59.8,"apparentTemperatureMaxTime":1539723600},{"time":1539756000,"summary":"Mostly cloudy throughout the day.","icon":"partly-cloudy-day","sunriseTime":1539785191,"sunsetTime":1539822856,"moonPhase":0.28,"precipIntensity":0,"precipIntensityMax":0.0001,"precipIntensityMaxTime":1539810000,"precipProbability":0,"temperatureHigh":65.14,"temperatureHighTime":1539810000,"temperatureLow":42.87,"temperatureLowTime":1539871200,"apparentTemperatureHigh":65.14,"apparentTemperatureHighTime":1539810000,"apparentTemperatureLow":37.7,"apparentTemperatureLowTime":1539871200,"dewPoint":33.65,"humidity":0.52,"pressure":1018.58,"windSpeed":5.89,"windGust":16.38,"windGustTime":1539766800,"windBearing":242,"cloudCover":0.36,"uvIndex":2,"uvIndexTime":1539799200,"visibility":10,"ozone":276.88,"temperatureMin":40.59,"temperatureMinTime":1539784800,"temperatureMax":65.14,"temperatureMaxTime":1539810000,"apparentTemperatureMin":36.1,"apparentTemperatureMinTime":1539784800,"apparentTemperatureMax":65.14,"apparentTemperatureMaxTime":1539810000},{"time":1539842400,"summary":"Mostly cloudy until evening.","icon":"partly-cloudy-day","sunriseTime":1539871703,"sunsetTime":1539909121,"moonPhase":0.31,"precipIntensity":0,"precipIntensityMax":0.0002,"precipIntensityMaxTime":1539885600,"precipProbability":0,"temperatureHigh":57.17,"temperatureHighTime":1539896400,"temperatureLow":37.99,"temperatureLowTime":1539954000,"apparentTemperatureHigh":57.17,"apparentTemperatureHighTime":1539896400,"apparentTemperatureLow":33.69,"apparentTemperatureLowTime":1539957600,"dewPoint":33.63,"humidity":0.55,"pressure":1020.58,"windSpeed":7.48,"windGust":20.53,"windGustTime":1539856800,"windBearing":286,"cloudCover":0.61,"uvIndex":2,"uvIndexTime":1539889200,"visibility":10,"ozone":271.42,"temperatureMin":42.87,"temperatureMinTime":1539871200,"temperatureMax":57.17,"temperatureMaxTime":1539896400,"apparentTemperatureMin":37.7,"apparentTemperatureMinTime":1539871200,"apparentTemperatureMax":57.17,"apparentTemperatureMaxTime":1539896400},{"time":1539928800,"summary":"Partly cloudy starting in the afternoon.","icon":"partly-cloudy-night","sunriseTime":1539958216,"sunsetTime":1539995387,"moonPhase":0.34,"precipIntensity":0,"precipIntensityMax":0.0002,"precipIntensityMaxTime":1539972000,"precipProbability":0,"temperatureHigh":60.43,"temperatureHighTime":1539986400,"temperatureLow":39.85,"temperatureLowTime":1540044000,"apparentTemperatureHigh":60.43,"apparentTemperatureHighTime":1539986400,"apparentTemperatureLow":36.6,"apparentTemperatureLowTime":1540044000,"dewPoint":33.17,"humidity":0.56,"pressure":1023.76,"windSpeed":6.13,"windGust":32.71,"windGustTime":1540004400,"windBearing":199,"cloudCover":0.19,"uvIndex":2,"uvIndexTime":1539972000,"visibility":10,"ozone":264.49,"temperatureMin":37.99,"temperatureMinTime":1539954000,"temperatureMax":60.43,"temperatureMaxTime":1539986400,"apparentTemperatureMin":33.69,"apparentTemperatureMinTime":1539957600,"apparentTemperatureMax":60.43,"apparentTemperatureMaxTime":1539986400},{"time":1540015200,"summary":"Partly cloudy until afternoon.","icon":"partly-cloudy-day","sunriseTime":1540044729,"sunsetTime":1540081653,"moonPhase":0.37,"precipIntensity":0,"precipIntensityMax":0.0002,"precipIntensityMaxTime":1540058400,"precipProbability":0,"temperatureHigh":64.23,"temperatureHighTime":1540069200,"temperatureLow":37.5,"temperatureLowTime":1540130400,"apparentTemperatureHigh":64.23,"apparentTemperatureHighTime":1540069200,"apparentTemperatureLow":33.46,"apparentTemperatureLowTime":1540130400,"dewPoint":35.85,"humidity":0.57,"pressure":1017.91,"windSpeed":1.16,"windGust":29.02,"windGustTime":1540015200,"windBearing":217,"cloudCover":0.32,"uvIndex":2,"uvIndexTime":1540062000,"visibility":10,"ozone":272.22,"temperatureMin":39.85,"temperatureMinTime":1540044000,"temperatureMax":64.23,"temperatureMaxTime":1540069200,"apparentTemperatureMin":36.6,"apparentTemperatureMinTime":1540044000,"apparentTemperatureMax":64.23,"apparentTemperatureMaxTime":1540069200}]},"flags":{"sources":["ecpa","cmc","gfs","icon","isd","madis","nam","sref"],"nearest-station":1.562,"units":"us"},"offset":-6}`
const expectedTemp = 32.94
const expectedSummary = "Mostly Cloudy"

func TestForecast(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	ds := NewDarkSky(ts.URL, "secret")
	w, err := ds.Forecast(Location{})
	if err != nil {
		t.Fatal(err)
	}
	if w.Currently.Temperature != expectedTemp {
		t.Errorf("Expected temperature %f, got %f.", expectedTemp, w.Currently.Temperature)
	}
	if w.Currently.Summary != expectedSummary {
		t.Errorf("Expected summary %q, got %q.", expectedSummary, w.Currently.Summary)
	}
}

func TestBadSecret(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "permission denied")
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	ds := NewDarkSky(ts.URL, "secret")
	_, err := ds.Forecast(Location{})

	if err != ErrUnauthorized {
		t.Errorf("Expected ErrUnauthorized, got %v.", err)
	}
}
