import 'package:holovn_mobile/models/channel.dart';
import 'package:json_annotation/json_annotation.dart';
part 'schedule.g.dart';

@JsonSerializable(explicitToJson: true)
class Schedule{
  final String id;
  final String title;
  final String type;
  final DateTime publishedAt;
  final DateTime availableAt;
  final DateTime startScheduled;
  final int duration;
  final String status;
  final Channel channel;

  Schedule(this.id, this.title, this.type, this.publishedAt, this.availableAt, this.startScheduled, this.duration, this.status, this.channel);
  factory Schedule.fromJson(Map<String, dynamic> json) => _$ScheduleFromJson(json);
  Map<String, dynamic> toJson() => _$ScheduleToJson(this);
}