from mitmproxy import ctx

ctx.master.addons.remove(
   *[i for i in ctx.master.addons.chain if i.__class__.__name__ == 'DisableH2C']
)
