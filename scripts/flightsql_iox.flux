 import "contrib/qxip/iox"
 iox.from(
     bucket: "test",
     host: "eu-central-1-1.aws.cloud2.influxdata.com:443",
     secure: "true",
     token: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXw==",
     limit: "10",
     table: "logs",
     start: -12h,
 )
