terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Availability Zone
provider "aws" {
  region = "ap-southeast-2"
}

resource "aws_scheduler_schedule_group" "precise-schedule-group" {
  name = "precise_schedule_group"
}

# SQS Queue
resource "aws_sqs_queue" "target-one" {
  name = "target-one"
  delay_seconds             = 90
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 10
  redrive_policy = jsonencode({
    deadLetterTargetArn = aws_sqs_queue.target-one-dlq.arn
    maxReceiveCount     = 4
  })
}

# Dead Letter Queue
resource "aws_sqs_queue" "target-one-dlq" {
  name = "target-one-dlq"
}

resource "aws_sqs_queue_redrive_allow_policy" "terraform_queue_redrive_allow_policy" {
  queue_url = aws_sqs_queue.target-one.id

  redrive_allow_policy = jsonencode({
    redrivePermission = "byQueue",
    sourceQueueArns   = [aws_sqs_queue.target-one.arn]
  })
}