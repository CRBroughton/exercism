using System;

static class LogLine
{
    public static string Message(string logLine)
    {
        var result = logLine switch
        {
            string line when line.Contains("[INFO]") => logLine["[INFO]: ".Length..].Trim(),
            string line when line.Contains("[WARNING]") => logLine["[WARNING]: ".Length..].Trim(),
            string line when line.Contains("[ERROR]") => logLine["[ERROR]: ".Length..].Trim(),
            _ => throw new NotImplementedException(),
        };
        return result;
    }

    public static string LogLevel(string logLine)
    {
        var result = logLine switch
        {
            string line when line.Contains("[INFO]") => "info",
            string line when line.Contains("[WARNING]") => "warning",
            string line when line.Contains("[ERROR]") =>  "error",
            _ => throw new NotImplementedException(),
        };
        return result;
    }

    public static string Reformat(string logLine)
    {
        var result = logLine switch
        {
            string line when line.Contains("[INFO]") => $"{logLine["[INFO]: ".Length..].Trim() } (info)",
            string line when line.Contains("[WARNING]") => $"{logLine["[WARNING]: ".Length..].Trim() } (warning)",
            string line when line.Contains("[ERROR]") => $"{logLine["[ERROR]: ".Length..].Trim() } (error)",
            _ => throw new NotImplementedException(),
        };
        return result;
    }
}
