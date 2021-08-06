import 'package:json_annotation/json_annotation.dart';
part 'channel.g.dart';

@JsonSerializable()
class Channel{
  final String channelId;
  final String name;
  final String org;
  final String type;
  final String photo;
  final String englishName;

  Channel(this.channelId, this.name, this.org, this.type, this.photo, this.englishName);
  factory Channel.fromJson(Map<String, dynamic> json) => _$ChannelFromJson(json);
  Map<String, dynamic> toJson() => _$ChannelToJson(this);
}