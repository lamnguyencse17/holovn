import 'package:holovn_mobile/models/channel.dart';
import 'package:json_annotation/json_annotation.dart';
part 'schedule.g.dart';

@JsonSerializable(explicitToJson: true)
class Schedule{
  final String scheduleId;
  final String title;
  final String type;
  final String publishedAt;
  final String availableAt;
  final int duration;
  final String status;
  final Channel channel;

  Schedule(this.scheduleId, this.title, this.type, this.publishedAt, this.availableAt, this.duration, this.status, this.channel);
  factory Schedule.fromJson(Map<String, dynamic> json) => _$ScheduleFromJson(json);
  Map<String, dynamic> toJson() => _$ScheduleToJson(this);
}