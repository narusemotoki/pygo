#!/usr/bin/env python
# -*- coding: utf-8 -*-
import falcon
import ujson as json

class Count(object):
    def on_post(self, req, resp):
        resp.body = json.dumps({
            'count': len(json.loads(req.stream.read())['key'].split(' '))
        })

application = falcon.API()
application.add_route('/count', Count())
