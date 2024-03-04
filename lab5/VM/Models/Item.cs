using System.ComponentModel.DataAnnotations;

namespace VM.Models;

public class Item
{
    public int Id { get; set; }
    public ItemType? ItemType { get; set; }
    [Range(0, 100)]
    public int Quality { get; set; }
    public Player? Owner { get; set; }
}
