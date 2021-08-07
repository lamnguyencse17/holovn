import 'package:flutter/material.dart';
import 'package:holovn_mobile/router/route_path.dart';

class RouteParser extends RouteInformationParser<RoutePath> {
  @override
  Future<RoutePath> parseRouteInformation(
      RouteInformation routeInformation) async {
    final uri = Uri.parse(routeInformation.location!);
    // Handle '/'
    if (uri.pathSegments.length == 0) {
      return RoutePath.home();
    }

    if (uri.pathSegments.length == 2) {
      if (uri.pathSegments[0] != 'live') return RoutePath.unknown();
      var liveId = uri.pathSegments[1];
      // if (liveId == null) return RoutePath.unknown();
      return RoutePath.live(null, liveId);
    }

    // Handle unknown routes
    return RoutePath.unknown();
  }

  @override
  RouteInformation restoreRouteInformation(RoutePath path) {
    if (path.isHomePage) {
      return RouteInformation(location: '/');
    }
    if (path.isLivePage) {
      return RouteInformation(location: '/live/${path.liveId}');
    }
    return RouteInformation(location: '/404');
  }
}
