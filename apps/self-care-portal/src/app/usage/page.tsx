import React from 'react';

export default function UsageVisualization() {
  return (
    <div className="p-8 space-y-6">
      <header>
        <h1 className="text-2xl font-bold">Data Usage Analysis</h1>
        <p className="text-slate-500">Real-time view of your consumption from ClickHouse</p>
      </header>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className="bg-white p-6 rounded-xl border shadow-sm h-64">
          <h3 className="font-semibold mb-4">Daily Consumption (GB)</h3>
          <div className="flex items-center justify-center h-full text-slate-300 italic">
            [Usage Line Chart: last 30 days]
          </div>
        </div>

        <div className="bg-white p-6 rounded-xl border shadow-sm">
          <h3 className="font-semibold mb-4">Top Applications</h3>
          <ul className="space-y-4">
            <UsageItem app="Streaming" usage="145 GB" percent="65" color="bg-blue-500" />
            <UsageItem app="Social Media" usage="45 GB" percent="20" color="bg-pink-500" />
            <UsageItem app="Work/VPN" usage="22 GB" percent="10" color="bg-green-500" />
          </ul>
        </div>
      </div>
    </div>
  );
}

function UsageItem({ app, usage, percent, color }: any) {
  return (
    <div className="space-y-1">
      <div className="flex justify-between text-sm font-medium">
        <span>{app}</span>
        <span>{usage}</span>
      </div>
      <div className="w-full h-2 bg-slate-100 rounded-full overflow-hidden">
        <div className={\`h-full \${color}\`} style={{ width: \`\${percent}%\` }}></div>
      </div>
    </div>
  );
}
