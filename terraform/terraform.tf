variable "mb_key" {}

module "machinebox" {
  source = "github.com/nicholasjackson/terraform-modules/elasticbeanstalk-docker"

  instance_type = "m4.4xlarge"

  application_name        = "gopherdataday"
  application_description = "Machinebox server"
  application_environment = "development"
  application_version     = "1.0.0"
  docker_image            = "machinebox/textbox"
  docker_tag              = "latest"
  docker_ports            = ["8080"]
  health_check            = "/info"
  env_vars                = ["MB_KEY", "${var.mb_key}"]
  elb_scheme              = "external"
}

data "aws_route53_zone" "selected" {
  name = "demo.gs."
}

resource "aws_route53_record" "textbox" {
  zone_id = "${data.aws_route53_zone.selected.zone_id}"
  name    = "textbox.${data.aws_route53_zone.selected.name}"
  type    = "CNAME"
  ttl     = "300"
  records = ["${module.machinebox.cname}"]
}
