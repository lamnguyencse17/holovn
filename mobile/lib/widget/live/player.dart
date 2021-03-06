import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart' show kIsWeb;
import 'package:holovn_mobile/widget/live/youtube_mobile_player.dart';
import 'package:holovn_mobile/widget/live/youtube_web_player.dart';
class Player extends StatelessWidget {
  final String liveId;
  final String status;
  Player(this.liveId, this.status);

  @override
  Widget build(BuildContext context) {
    if (kIsWeb){
      return YoutubeWebPlayer(liveId: liveId);
    }
    return YoutubeMobilePlayer(liveId: liveId, status: status,);
  }
}
