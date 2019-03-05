import http.client as hc


class HttpClient(object):
    _conn = None

    def __init__(self, host):
        self._conn = hc.HTTPConnection(host=host)

    def request(self, method, url, data=None, headers={}):
        self._conn.request(method, url, data, headers)
        resp = self._conn.getresponse()
        b = resp.read()
        data = str(b, encoding='utf-8')
        return data

